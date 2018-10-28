package configuration

import (
	"encoding/json"
	"errors"
	//	. "github.com/aeternas/SwadeshNess/endpoints"
	. "github.com/aeternas/SwadeshNess/language"
	. "github.com/aeternas/SwadeshNess/wrappers"
	"log"
	"os"
)

var (
	osWrapper AnyOsWrapper
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
	Path string
}

func init() {
	osWrapper = new(OsWrapper)
}

func (r *Reader) ReadConfiguration() (Configuration, error) {
	var p string = (*r).Path
	file, _ := os.Open(p)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	apiKey := osWrapper.GetEnv("YANDEX_API_KEY")
	var translationEndpoint string = osWrapper.GetEnv("TRANSLATION_ENDPOINT")
	var groupEndpoint string = osWrapper.GetEnv("GROUP_ENDPOINT")
	configuration.ApiKey = apiKey
	endpoints := Endpoints{TranslationEndpoint: translationEndpoint, GroupsEndpoint: groupEndpoint}
	configuration.EEndpoints = endpoints
	if err != nil {
		log.Fatal("Unable to read database")
		return Configuration{}, errors.New("Failed to read database")
	}
	return configuration, nil
}
