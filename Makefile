.PHONY: build test clean install

build:
	go build -o bin/leetx ./cmd/leetx

install:
	go install ./cmd/leetx

test:
	go test -v ./...

clean:
	go clean
	rm -rf bin/
