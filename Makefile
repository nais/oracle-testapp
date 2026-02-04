.PHONY: build run clean test

build:
	go build -o oracle-testapp .

run: build
	./oracle-testapp

test:
	go test -v ./...

clean:
	rm -f oracle-testapp
