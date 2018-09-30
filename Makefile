all: build

build:
	go build -o bin/fix-corrupt-zsh main.go

install:
	go install

clean:
	rm -rf bin
