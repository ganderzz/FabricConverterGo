package utils

import (
	"encoding/json"
	"os"

	"github.com/ganderzz/FabricConverterGo/src/fabric"
)

// ConvertFileToFabricJSON Convert file to JSON object
func ConvertFileToFabricJSON(file *os.File) (*fabric.FabricBaseObject, error) {
	fabricObj := &fabric.FabricBaseObject{}
	err := json.NewDecoder(file).Decode(fabricObj)
	if err != nil {
		return nil, err
	}

	return fabricObj, nil
}

// ConvertBytesToFabricJSON Convert byte array to JSON object
func ConvertBytesToFabricJSON(data []byte) (*fabric.FabricBaseObject, error) {
	fabricObj := &fabric.FabricBaseObject{}

	err := json.Unmarshal(data, fabricObj)
	if err != nil {
		return nil, err
	}

	return fabricObj, nil
}

// GetFabricJSONFromFile Loads JSON and converts it to a fabric object
func GetFabricJSONFromFile(path string) (*fabric.FabricBaseObject, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	return ConvertFileToFabricJSON(file)
}
