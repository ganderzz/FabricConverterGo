all:
		go run ./src/*.go

test:
		go run ./src/*.go input/fabric.json out.png

build:
		go build -o fti ./src/*.go

clean:
	rm out.png fti