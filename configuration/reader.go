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
	var translationEndpoint string = lReader.OsWrapper.GetEnvFallback(TRANSLATION_ENDPOINT, "/")
	var groupEndpoint string = lReader.OsWrapper.GetEnvFallback(GROUP_ENDPOINT, "/groups")
	var versionEndpoint string = lReader.OsWrapper.GetEnvFallback(VERSION_ENDPOINT, "/version")
	var version string = lReader.OsWrapper.GetEnvFallback(VERSION, "0")
	var serverKeyPath string = lReader.OsWrapper.GetEnvFallback(SERVER_KEY, "certs/server.key")
	var serverCertPath string = lReader.OsWrapper.GetEnvFallback(SERVER_CERT, "certs/server.crt")
	configuration.ApiKey = apiKey
	configuration.Version = version
	endpoints := Endpoints{TranslationEndpoint: translationEndpoint, GroupsEndpoint: groupEndpoint, VersionEndpoint: versionEndpoint}
	security := Security{ServerKeyPath: serverKeyPath, ServerCertPath: serverCertPath}
	configuration.EEndpoints = endpoints
	configuration.Security = security
	if err != nil {
		log.Printf("Configuration decoding failed")
		return Configuration{}, errors.New("Failed to read database")
	}
	return configuration, nil
}
