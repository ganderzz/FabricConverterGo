package server

import (
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ganderzz/FabricConverterGo/src/utils"

	"github.com/fogleman/gg"
)

func writeImage(w http.ResponseWriter, img *image.Image) {
	if err := png.Encode(w, *img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
}

//HandleUploadController Handle upload
func HandleUploadController(writer http.ResponseWriter, reader *http.Request) {
	body, err := ioutil.ReadAll(reader.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(writer, "Can't read body", http.StatusBadRequest)
		return
	}

	json, err := utils.ConvertBytesToFabricJSON(body)
	if err != nil {
		log.Printf("Error converting to JSON: %v", err)
		http.Error(writer, "Can't convert Fabric to JSON object.", http.StatusBadRequest)
		return
	}

	context := gg.NewContext(800, 800)
	for _, o := range json.Objects {
		o.Parse(context)
	}

	img := context.Image()

	writeImage(writer, &img)
}
