CC = g++
AR = ar cr
RM = rm -f

LIB = digidrum.o YmLoad.o Ym2149Ex.o YmMusic.o YmUserInterface.o

CFLAGS = -Wall -Werror -Wextra -pedantic -std=c++17
CFLAGSLIGHT = -Wall -Wextra

libym.a: $(LIB)
		$(AR) libstsound.a $(LIB)

digidrum.o: digidrum.cpp YmTypes.h
		$(CC) $(CFLAGS) -c digidrum.cpp

Ym2149Ex.o: Ym2149Ex.cpp Ym2149Ex.h YmTypes.h
		$(CC) $(CFLAGSLIGHT) -c Ym2149Ex.cpp

YmLoad.o: YmLoad.cpp YmMusic.h lzh/lzh.h
		$(CC) $(CFLAGSLIGHT) -c YmLoad.cpp

YmMusic.o: YmMusic.cpp YmMusic.h
		$(CC) $(CFLAGSLIGHT) -c YmMusic.cpp

YmUserInterface.o: YmUserInterface.cpp StSoundLibrary.h YmMusic.h
		$(CC) $(CFLAGS) -c YmUserInterface.cpp

.PHONY: clean
clean:
		$(RM) *~ *.o *.a
