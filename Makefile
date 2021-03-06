pkg=deconstructedgeohash

check: lint test

test:
	go test -v

lint:
	golint
	asmfmt -w *.s

%.h: %.py
	python $< > $@
