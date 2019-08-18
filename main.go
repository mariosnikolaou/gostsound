package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -L${SRCDIR}/StSoundLibrary -L${SRCDIR}/StSoundLibrary/LZH -lstsound -llzh -lc++
// #include <stdlib.h>
// #include "StSoundLibrary/StSoundLibrary.h"
import "C"
import (
	"flag"
	"fmt"
	"unsafe"

	"github.com/hajimehoshi/oto"
)

const (
	sampleRate      = 44100
	channelNum      = 1
	bitDepthInBytes = 2
	sampleWindow    = 1024
)

func main() {
	flag.Parse()
	args := flag.Args()

	filename := args[0]
	fmt.Println("Reading: ", filename)

	pMusic := C.ymMusicCreate()
	defer C.free(unsafe.Pointer(pMusic))

	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	C.ymMusicLoad(pMusic, cfilename)

	info := C.ymMusicInfo_t{}
	C.ymMusicGetInfo(pMusic, &info)

	fmt.Println("SongName:       ", C.GoString(info.pSongName))
	fmt.Println("SongAuthor:     ", C.GoString(info.pSongAuthor))
	fmt.Println("SongComment:    ", C.GoString(info.pSongComment))
	fmt.Println("SongType:       ", C.GoString(info.pSongType))
	fmt.Println("SongPlayer:     ", C.GoString(info.pSongPlayer))
	fmt.Println("musicTimeInSec: ", info.musicTimeInSec)
	fmt.Println("musicTimeInMs:  ", info.musicTimeInMs)

	C.ymMusicSetLoopMode(pMusic, C.YMFALSE)
	C.ymMusicStop(pMusic)
	C.ymMusicPlay(pMusic)

	// Allocate buffer (little-endian 16 bit, mono 44100 Hz)
	buf := C.malloc(C.sizeof_ymsample * sampleWindow * 2)
	defer C.free(unsafe.Pointer(buf))

	// Create the player
	ctx, err := oto.NewContext(sampleRate, channelNum, bitDepthInBytes, 4096)
	if err != nil {
		panic(err)
	}
	p := ctx.NewPlayer()

	// Compute next buffer and copy to player (blocking)
	for {
		done := C.ymMusicCompute(pMusic, (*C.ymsample)(buf), sampleWindow)
		if done == C.YMFALSE {
			break
		}
		_, err = p.Write(C.GoBytes(buf, sampleWindow*2))
		if err != nil {
			panic(err)
		}
	}

	err = p.Close()
	if err != nil {
		panic(err)
	}

}
