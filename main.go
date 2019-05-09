package main

import (
	. "github.com/aeternas/SwadeshNess-packages/language"
	ApiClient "github.com/aeternas/SwadeshNess/apiClient"
	ClientMiddlewares "github.com/aeternas/SwadeshNess/clientMiddlewares"
	Config "github.com/aeternas/SwadeshNess/configuration"
	. "github.com/aeternas/SwadeshNess/handlers"
	HTTPApiClient "github.com/aeternas/SwadeshNess/httpApiClient"
	ServerMiddlewares "github.com/aeternas/SwadeshNess/serverMiddlewares"
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"log"
	"net/http"
	"time"
)

var (
	languageGroups     []LanguageGroup
	reader             Config.AnyReader
	translationHandler AnyHandler
	groupListHandler   AnyHandler
	versionHandler     AnyHandler
	configuration      Config.Configuration
	apiClient          ApiClient.ApiClient
)

const (
	DATABASE_CONFIG = "configuration/db.json"
)

func init() {
	var wrapper = Wrappers.New(new(Wrappers.OsWrapper))
	var lReader *Config.Reader = &Config.Reader{Path: DATABASE_CONFIG, OsWrapper: wrapper}

	reader = lReader
	lConfiguration, _ := reader.ReadConfiguration()
	configuration = lConfiguration

	clientMiddlewares := []ClientMiddlewares.ClientMiddleware{
		ClientMiddlewares.NewCachingDefaultClientMiddleware(&configuration),
		ClientMiddlewares.NewDefaultClientMiddleware(),
		ClientMiddlewares.NewAuthClientMiddleware(configuration.ApiKey),
		ClientMiddlewares.NewLoggerClientMiddleware(),
	}

	httpApiClient := &HTTPApiClient.HTTPApiClient{Client: &http.Client{Timeout: 10 * time.Second}, Middlewares: clientMiddlewares}
	apiClient = httpApiClient

	csm := ServerMiddlewares.NewCachingDefaultServerMiddleware(&configuration)
	lsm := ServerMiddlewares.NewLoggerServerMiddleware()
	translationHandler = &TranslationHandler{
		Config:            &configuration,
		ServerMiddlewares: []ServerMiddlewares.ServerMiddleware{csm},
		ApiClient:         apiClient,
	}

	groupListHandler = &GroupListHandler{Config: &configuration, ServerMiddlewares: []ServerMiddlewares.ServerMiddleware{lsm}}
	versionHandler = &VersionHandler{Config: &configuration}
}

func main() {
	languageGroups = configuration.Languages
	http.HandleFunc(configuration.EEndpoints.GroupsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		groupListHandler.HandleRequest(w, r)
	})
	http.HandleFunc(configuration.EEndpoints.TranslationEndpoint, func(w http.ResponseWriter, r *http.Request) {
		translationHandler.HandleRequest(w, r)
	})
	http.HandleFunc(configuration.EEndpoints.VersionEndpoint, func(w http.ResponseWriter, r *http.Request) {
		versionHandler.HandleRequest(w, r)
	})
	if configuration.Security.NeedsHTTPS {
		log.Fatal(http.ListenAndServeTLS(":8080", configuration.Security.ServerCertPath, configuration.Security.ServerKeyPath, nil))
	} else {
		log.Fatal(http.ListenAndServe(":8081", nil))
	}
}
