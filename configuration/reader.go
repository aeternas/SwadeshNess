package configuration

import (
	"encoding/json"
	"errors"
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"log"
)

type AnyReader interface {
	ReadConfiguration() (Configuration, error)
}

type Reader struct {
	Path      string
	OsWrapper Wrappers.AnyOsWrapper
}

func (r *Reader) ReadConfiguration() (Configuration, error) {
	lReader := *r
	var p string = lReader.Path
	fileOpened := lReader.OsWrapper.Open(p)
	file := fileOpened.F
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	apiKey := lReader.OsWrapper.GetEnv(API_KEY)
	var translationEndpoint string = lReader.OsWrapper.GetEnv(TRANSLATION_ENDPOINT)
	var groupEndpoint string = lReader.OsWrapper.GetEnv(GROUP_ENDPOINT)
	configuration.ApiKey = apiKey
	endpoints := Endpoints{TranslationEndpoint: translationEndpoint, GroupsEndpoint: groupEndpoint}
	configuration.EEndpoints = endpoints
	if err != nil {
		log.Printf("Configuration decoding failed")
		return Configuration{}, errors.New("Failed to read database")
	}
	return configuration, nil
}
