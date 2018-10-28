package configuration

import (
	"encoding/json"
	"errors"
	. "github.com/aeternas/SwadeshNess/language"
	. "github.com/aeternas/SwadeshNess/wrappers"
	"log"
	"os"
)

type Configuration struct {
	Languages  []LanguageGroup
	ApiKey     string
	Credits    string
	EEndpoints Endpoints
}

type AnyReader interface {
	ReadConfiguration() (Configuration, error)
}

type Reader struct {
	Path      string
	OsWrapper AnyOsWrapper
}

func (r *Reader) ReadConfiguration() (Configuration, error) {
	lReader := *r
	var p string = lReader.Path
	file, _ := os.Open(p)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	apiKey := lReader.OsWrapper.GetEnv("YANDEX_API_KEY")
	var translationEndpoint string = lReader.OsWrapper.GetEnv("TRANSLATION_ENDPOINT")
	var groupEndpoint string = lReader.OsWrapper.GetEnv("GROUP_ENDPOINT")
	configuration.ApiKey = apiKey
	endpoints := Endpoints{TranslationEndpoint: translationEndpoint, GroupsEndpoint: groupEndpoint}
	configuration.EEndpoints = endpoints
	if err != nil {
		log.Fatal("Unable to read database")
		return Configuration{}, errors.New("Failed to read database")
	}
	return configuration, nil
}
