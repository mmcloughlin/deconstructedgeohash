CC=clang
CFLAGS=-std=c11 -O3 -march=native -Wextra
ASMFLAGS=-S -masm=intel -mno-red-zone -mstackrealign -mllvm \
				 -inline-threshold=1000 -fno-asynchronous-unwind-tables \
				 -fno-exceptions -fno-rtti

lib=geohash
test=test_$(lib)
bench=bench_$(lib)

srcs = $(wildcard *.c)
deps = $(srcs:.c=.d)

all: $(test) $(bench) $(lib).c.s

$(test): $(test).o $(lib).o
$(bench): $(bench).o $(lib).o

$(test).o $(bench).o: testvector.h
$(bench).o: benchmark.h

%.o: %.c
	$(CC) -c -MMD -MP $(CFLAGS) $< -o $@

-include $(deps)

%.h: %.go godeps
	go run $< > $@

%.c.s: %.c
	$(CC) $(CFLAGS) $(ASMFLAGS) $^ -o $@

benchmark.h:
	wget -O $@ https://raw.githubusercontent.com/lemire/Code-used-on-Daniel-Lemire-s-blog/e96243e014/tmp/benchmark.h

godeps:
	go get github.com/mmcloughlin/deconstructedgeohash
	go get github.com/mmcloughlin/spherand

clean:
	$(RM) $$(cat .gitignore)

.PHONY: all clean run tools
