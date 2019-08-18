CC = gcc
AR = ar cr
RM = rm -f

CFLAGS = -Wall -Werror -Wextra -pedantic -std=c++17

liblzh.a: lzh.o
		$(AR) liblzh.a lzh.o

lzh.o: lzh.cpp lzh.h
		$(CC) $(CFLAGS) -c lzh.cpp

.PHONY: clean
clean:
		$(RM) *~ *.o *.a
