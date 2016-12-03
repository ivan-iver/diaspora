SHELL := bash
# VERSION := 'API version\:$(shell date -u +%Y%m%d.%H%M%S)\($(shell git rev-parse --short HEAD)\)'
VERSION := ($(shell git rev-parse --short HEAD))

export VERSION;

build:
	go build -v -ldflags "-X github.com/ivan-iver/diaspora/lib.hash=${VERSION}" -o bin/diaspora github.com/ivan-iver/diaspora
	@cp templates/db.conf bin/

clean:
	rm -r bin/*

init:
	mkdir -p bin log
	go get -v -u ./...

