.PHONY: all test build clean

all: clean test build

build: 
	mkdir -p build
	go build -o build ./...

test:
	go test -v -coverprofile=tests/results/cover.out ./...

cover:
	go tool cover -html=tests/results/cover.out -o tests/results/cover.html

clean:
	rm -rf build/*
	go clean ./...

container:
	podman build -t  quay.io/luigizuccarelli/golang-cicd-webconsole:1.15.6 .

push:
	podman push --authfile=/home/lzuccarelli/confif.json quay.io/luigizuccarelli/golang-cicd-webconsole:1.15.6 
