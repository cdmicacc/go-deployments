default: build

all: build test

clean:
	goop exec go clean

build: deps
	goop exec go build

deps:
	goop install

test: deps
	goop exec go test ./...

