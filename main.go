package main

import (
	. "github.com/aeternas/SwadeshNess-packages/language"
	Config "github.com/aeternas/SwadeshNess/configuration"
	. "github.com/aeternas/SwadeshNess/handlers"
	ServerMiddlewares "github.com/aeternas/SwadeshNess/serverMiddlewares"
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"log"
	"net/http"
)

var (
	languageGroups     []LanguageGroup
	reader             Config.AnyReader
	translationHandler AnyHandler
	groupListHandler   AnyHandler
	versionHandler     AnyHandler
	configuration      Config.Configuration
)

func init() {
	var wrapper = Wrappers.New(new(Wrappers.OsWrapper))
	var lReader *Config.Reader = &Config.Reader{Path: "configuration/db.json", OsWrapper: wrapper}
	reader = lReader
	lConfiguration, _ := reader.ReadConfiguration()
	configuration = lConfiguration
	cm := ServerMiddlewares.NewCachingDefaultServerMiddleware()
	translationHandler = &TranslationHandler{Config: &configuration, Middlewares: []ServerMiddlewares.ServerMiddleware{cm}}
	groupListHandler = &GroupListHandler{Config: &configuration}
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
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
