package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -L${SRCDIR}/StSoundLibrary -L${SRCDIR}/StSoundLibrary/lzh -lstsound -llzh -lc++
// #include <stdlib.h>
// #include "StSoundLibrary/StSoundLibrary.h"
import "C"

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
)

const (
	sampleRate   = 44100
	channelNum   = 1
	sampleWindow = 1024
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: gostsound <filename>")
		return
	}

	filename := args[0]
	fmt.Println("Reading: ", filename)

	pMusic := C.ymMusicCreate()
	if pMusic == nil {
		log.Fatal("ymMusicCreate failed")
	}
	defer C.ymMusicDestroy(pMusic)

	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Could not read file %s: %v", filename, err)
	}
	fmt.Printf("Size of data: %d\n", len(data))

	cdata := C.CBytes(data)
	defer C.free(cdata)

	if C.ymMusicLoadMemory(pMusic, cdata, C.uint(len(data))) == C.YMFALSE {
		errMsg := C.GoString(C.ymMusicGetLastError(pMusic))
		log.Fatalf("ymMusicLoadMemory failed: %s", errMsg)
	}

	// Get and print music info
	info := C.ymMusicInfo_t{}
	C.ymMusicGetInfo(pMusic, &info)

	fmt.Println("SongName:       ", C.GoString(info.pSongName))
	fmt.Println("SongAuthor:     ", C.GoString(info.pSongAuthor))
	fmt.Println("SongComment:    ", C.GoString(info.pSongComment))
	fmt.Println("SongType:       ", C.GoString(info.pSongType))
	fmt.Println("SongPlayer:     ", C.GoString(info.pSongPlayer))
	fmt.Println("musicTimeInSec: ", info.musicTimeInSec)
	fmt.Println("musicTimeInMs:  ", info.musicTimeInMs)

	// Setup music playback
	C.ymMusicSetLoopMode(pMusic, C.YMFALSE)
	C.ymMusicStop(pMusic)
	C.ymMusicPlay(pMusic)

	// Allocate buffer (little-endian 16 bit, mono 44100 Hz)
	buf := C.malloc(C.sizeof_ymsample * sampleWindow * 2)
	defer C.free(buf)

	// Create the context and player
	op := &oto.NewContextOptions{
		SampleRate:   sampleRate,
		ChannelCount: channelNum,
		Format:       oto.FormatSignedInt16LE,
	}
	ctx, ready, err := oto.NewContext(op)
	if err != nil {
		log.Fatalf("oto.NewContext failed: %v", err)
	}
	<-ready

	r, w := io.Pipe()

	p := ctx.NewPlayer(r)
	p.Play()

	var playerErr error
	go func() {
		defer w.Close()
		for {
			done := C.ymMusicCompute(pMusic, (*C.ymsample)(buf), sampleWindow)
			if done == C.YMFALSE {
				break
			}
			_, err := w.Write(C.GoBytes(buf, sampleWindow*2))
			if err != nil {
				playerErr = err
				return
			}
		}
	}()

	for p.IsPlaying() {
		time.Sleep(100 * time.Millisecond)
	}

	if playerErr != nil {
		log.Fatalf("error during playback: %v", playerErr)
	}

	if err := p.Close(); err != nil {
		log.Fatalf("player.Close failed: %v", err)
	}
}
