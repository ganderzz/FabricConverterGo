package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"./server"
	"./utils"

	"github.com/fogleman/gg"
)

func main() {
	serve := flag.Bool("serve", false, "Start the HTTP server.")
	input := flag.String("input", "", "The input JSON file.")
	output := flag.String("output", "", "The output PNG file.")

	flag.Parse()

	if *serve {
		// Handle HTTP Server parsing
		http.HandleFunc("/", server.HandleUploadController)

		fmt.Println("Starting on: http://localhost:5000")

		err := http.ListenAndServe(":5000", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
			return
		}
	} else if len(*input) > 0 && len(*output) > 0 {
		// Handle command line parsing
		fabricObj, err := utils.GetFabricJSONFromFile(*input)
		if err != nil {
			fmt.Printf("Could not load JSON file (%s)\n", err.Error())
			return
		}

		//width, height := fabricObj.GetBounds()
		context := gg.NewContext(800, 800)
		for _, obj := range fabricObj.Objects {
			obj.Parse(context)
		}

		context.SavePNG(*output)
	} else {
		fmt.Println("Invalid arguments.")
		return
	}
}
