.PHONY: all clean test build run

all: clean test build run

clean:
	rm -f ./app 2>/dev/null

test:
	go test ./tests

build:
	go build -o ./app ./cmd

run:
	./app
