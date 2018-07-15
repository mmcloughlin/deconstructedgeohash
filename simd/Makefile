CC=clang
CFLAGS=-std=c11 -O3 -march=native # -Weverything

prog=geohash

all: $(prog)

run: $(prog)
	./$(prog)

clean:
	$(RM) $(prog)

.PHONY: all clean run
