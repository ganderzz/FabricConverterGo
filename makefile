all:
		go run ./src/*.go

test:
		go run ./src/*.go --input=input/fabric.json --output=out.png

build:
		go build -o fti ./src/*.go

clean:
	rm fti