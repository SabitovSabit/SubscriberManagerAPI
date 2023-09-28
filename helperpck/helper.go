package helper

import (
	"encoding/json"
	"os"
	"log"
	models "subscriptionApi/modelspck"
)
var File os.File
func GetValue() models.Config {
	File, err := os.Open("appsettings.json")
	var config models.Config

	decoder := json.NewDecoder(File)
	if err = decoder.Decode(&config); err != nil {
		log.Fatal(err)
	}
	return config
}
