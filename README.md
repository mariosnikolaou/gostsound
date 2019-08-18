GoSTSound

----

GoSTSound plays YM music files using the ST-Sound Library by Arnaud Carré.

The minimal Go wrapper uses cgo to call out to the ST-Sound libary to calculate
the PCM sound data.

To play the music GoSTSound uses the excellent Go library Oto
(https://github.com/hajimehoshi/oto).


## ST-Sound

The original ST-Sound library written in C/C++ is distributed under BSD
license. A copy of the original zip archive, StSound_1_43.zip, is included in
this repo.

ST-Sound library, Copyright (C) 1995-1999 Arnaud Carré
(http://leonard.oxg.free.fr)
