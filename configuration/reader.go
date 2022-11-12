package configuration

import (
	"encoding/json"
	"errors"
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"log"
	"os"
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
	apiKey := lReader.OsWrapper.GetEnvFallback(API_KEY, "")
	var translationEndpoint string = lReader.OsWrapper.GetEnvFallback(TRANSLATION_ENDPOINT, "/v1/")
	var groupEndpoint string = lReader.OsWrapper.GetEnvFallback(GROUP_ENDPOINT, "/v1/groups")
	var versionEndpoint string = lReader.OsWrapper.GetEnvFallback(VERSION_ENDPOINT, "/v1/version")
	var version string = lReader.OsWrapper.GetEnvFallback(VERSION, "0")
	var serverKeyPath string = lReader.OsWrapper.GetEnvFallback(SERVER_KEY, "certs/server.key")
	var serverCertPath string = lReader.OsWrapper.GetEnvFallback(SERVER_CERT, "certs/server.crt")
	var redisAddress string = lReader.OsWrapper.GetEnvFallback(REDIS_ADDRESS, "localhost")
	var needsHTTPS bool = false
	if args := os.Args; len(args) > 1 && os.Args[1] == "--https" {
		needsHTTPS = true
	}
	configuration.ApiKey = apiKey
	configuration.Version = version
	endpoints := Endpoints{TranslationEndpoint: translationEndpoint, GroupsEndpoint: groupEndpoint, VersionEndpoint: versionEndpoint, RedisAddress: redisAddress}
	security := Security{NeedsHTTPS: needsHTTPS, ServerKeyPath: serverKeyPath, ServerCertPath: serverCertPath}
	configuration.EEndpoints = endpoints
	configuration.Security = security
	if err != nil {
		log.Printf("Configuration decoding failed")
		return Configuration{}, errors.New("Failed to read database")
	}
	return configuration, nil
}
