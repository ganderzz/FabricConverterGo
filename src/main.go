package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fogleman/gg"
)

func printArgList() {
	fmt.Println("")
	fmt.Println("fti [inputFileName] [outputFileName]")
}

func getFabricJSONFromFile(path string) (*fabricBaseObject, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	fabricObj := &fabricBaseObject{}
	if json.NewDecoder(file).Decode(fabricObj) != nil {
		return nil, err
	}

	return fabricObj, nil
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Invalid arguments.")
		printArgList()
		return
	}

	fabricObj, err := getFabricJSONFromFile(args[0])
	if err != nil {
		fmt.Printf("Could not load JSON file (%s)\n", err.Error())
		return
	}

	width, height := fabricObj.GetBounds()
	context := gg.NewContext(int(width), int(height))
	for _, obj := range fabricObj.Objects {
		obj.Parse(context)
	}

	context.SavePNG(args[1])
}
