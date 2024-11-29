.PHONY: build test clean

build:
	go build -o bin/gobench ./cmd/gobench

test:
	go test -v ./...

clean:
	rm -rf bin

