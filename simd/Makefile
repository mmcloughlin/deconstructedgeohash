CC=clang
OPTFLAGS=-O3 -ffast-math -mllvm -x86-use-vzeroupper=0
CFLAGS=-std=c11 -mavx2 -Wextra $(OPTFLAGS)
IACAFLAGS=-DIACA -I$(IACAPATH)

lib=geohash
test=test_$(lib)
bench=bench_$(lib)

srcs = $(wildcard *.c)
deps = $(srcs:.c=.d)

all: $(test) $(bench) $(lib).c.s $(lib).objdump

$(test): $(test).o $(lib).o
$(bench): $(bench).o $(lib).o

$(test).o $(bench).o: testvector.h
$(bench).o: benchmark.h

%.o: %.c
	$(CC) -c -MMD -MP $(CFLAGS) $< -o $@

-include $(deps)

%.h: %.go
	go run $< > $@

%.c.s: %.c
	$(CC) $(CFLAGS) -S $^ -o $@

benchmark.h:
	wget -O $@ https://raw.githubusercontent.com/lemire/Code-used-on-Daniel-Lemire-s-blog/e96243e014/tmp/benchmark.h

%.iaca.o: %.c
	$(CC) -c $(IACAFLAGS) $(CFLAGS) $< -o $@

%.objdump: %.o
	objdump -disassemble-all $< > $@

clean:
	$(RM) $$(cat .gitignore)

.PHONY: all clean tools
