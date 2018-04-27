package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/golang-training/src/dino/dinowebportal"
)

type configuration struct {
	ServerAddress string `json:"webserver"` // (note: struct tags used to map fields)
}

func main() {
	file, err := os.Open("./src/dino/config.json")
	if err != nil {
		log.Fatal(err)
	}

	config := new(configuration)
	json.NewDecoder(file).Decode(config)

	dinowebportal.RunWebPortal(config.ServerAddress)
}
