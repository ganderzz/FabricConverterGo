package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"./server"
	"./utils"

	"github.com/fogleman/gg"
)

func printArgList() {
	fmt.Println("--- FTI ---")
	fmt.Println("HTTP:")
	fmt.Println("fti serve")
	fmt.Println("")
	fmt.Println("Command Line:")
	fmt.Println("fti [inputFileName] [outputFileName]")
	fmt.Println("")
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 && strings.ToLower(args[0]) == "serve" {
		// Handle HTTP Server parsing
		http.HandleFunc("/", server.HandleUploadController)

		fmt.Println("Starting on: http://localhost:5000")

		err := http.ListenAndServe(":5000", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
			return
		}
	} else if len(args) >= 2 {
		// Handle command line parsing
		fabricObj, err := utils.GetFabricJSONFromFile(args[0])
		if err != nil {
			fmt.Printf("Could not load JSON file (%s)\n", err.Error())
			return
		}

		//width, height := fabricObj.GetBounds()
		context := gg.NewContext(800, 800)
		for _, obj := range fabricObj.Objects {
			obj.Parse(context)
		}

		context.SavePNG(args[1])
	} else {
		fmt.Println("Invalid arguments.")
		printArgList()
		return
	}
}
