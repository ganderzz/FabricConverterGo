package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fogleman/gg"
)

func printArgList() {
	fmt.Println("ImageConverter")
	fmt.Println("------------------------")
	fmt.Println("imgcvt [inputFileName] [outputFileName]")
	fmt.Println("--extension   [png]")
}

func getFabricFile(path string) (*fabricBaseObject, error) {
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

	fabricObj, err := getFabricFile(args[0])
	if err != nil {
		fmt.Printf("Could not load JSON file (%s)\n", err.Error())
		return
	}

	dc := gg.NewContext(1000, 1000)
	for _, obj := range fabricObj.Objects {
		obj.Parse(dc)
	}

	dc.SavePNG(args[1])
}
