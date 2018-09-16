all:
		go run ./src/*.go

test:
		go run ./src/*.go input/fabric.json out.png

build:
		go build ./src/*.go