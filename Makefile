SHELL := bash
# VERSION := 'API version\:$(shell date -u +%Y%m%d.%H%M%S)\($(shell git rev-parse --short HEAD)\)'
VERSION := 'API:$(shell date -u +%Y%m%d.%H%M%S)-($(shell git rev-parse --short HEAD))'

export VERSION;

build:
	go build -v -ldflags "-X github.com/ivan-iver/diaspora/lib.version=${VERSION}" -o bin/diaspora github.com/ivan-iver/diaspora
#	@cp templates/app.conf bin/

start:
	./bin/api prod > ../logs/api.log 2>../logs/api.error.log &

clean:
	rm -r bin/*

init:
	mkdir -p bin log
	go get -v -u ./...

