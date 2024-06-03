# Makefile
.PHONY: all build run test clean

all: build

build:
	go build -o clock-app cmd/clock/main.go

run:
	./clock-app

test:
	go test -v ./...

clean:
	rm -f clock-app
