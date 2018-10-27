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

type AnyReader interface {
	ReadConfiguration() (Configuration, error)
}

type Reader struct {
	Path string
}

func (r *Reader) ReadConfiguration() (Configuration, error) {
	var p string = (*r).Path
	file, _ := os.Open(p)
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
