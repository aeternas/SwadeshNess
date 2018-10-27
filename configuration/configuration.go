package configuration

import (
	"encoding/json"
	"errors"
	. "github.com/aeternas/SwadeshNess/language"
	"log"
	"os"
)

type Configuration struct {
	Languages []LanguageGroup
}

func ReadConfiguration() (Configuration, error) {
	file, _ := os.Open("configuration/db.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("Unable to read database")
		return Configuration{}, errors.New("Failed to read database")
	}
	return configuration, nil
}
