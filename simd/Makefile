CC=clang
CFLAGS=-std=c11 -O3 -march=native # -Weverything

prog=geohash

all: $(prog)

$(prog): testvector.h

%: %.c
	$(CC) $(CFLAGS) $(filter %.c,$^) -o $@

%.h: %.go
	go run $< > $@

run: $(prog)
	./$(prog)

clean:
	$(RM) $(prog)

.PHONY: all clean run
