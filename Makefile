.PHONY: build run clean

build:
	go build -o oracle-testapp .

run: build
	./oracle-testapp

clean:
	rm -f oracle-testapp
