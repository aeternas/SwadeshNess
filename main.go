package main

import (
	. "github.com/aeternas/SwadeshNess-packages/language"
	Caching "github.com/aeternas/SwadeshNess/caching"
	Config "github.com/aeternas/SwadeshNess/configuration"
	. "github.com/aeternas/SwadeshNess/handlers"
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"log"
	"net/http"
)

var (
	languageGroups     []LanguageGroup
	reader             Config.AnyReader
	translationHandler AnyTranslationHandler
	groupListHandler   AnyGroupListHandler
	versionHandler     AnyVersionHandler
	configuration      Config.Configuration
	cw                 Caching.AnyCacheWrapper
)

func init() {
	var wrapper = Wrappers.New(new(Wrappers.OsWrapper))
	var lReader *Config.Reader = &Config.Reader{Path: "configuration/db.json", OsWrapper: wrapper}
	reader = lReader
	lConfiguration, _ := reader.ReadConfiguration()
	configuration = lConfiguration
	cw = Caching.NewRedisCachingWrapper()
	translationHandler = &TranslationHandler{Config: &configuration}
	groupListHandler = &GroupListHandler{Config: &configuration}
	versionHandler = &VersionHandler{Config: &configuration}
}

func main() {
	languageGroups = configuration.Languages
	http.HandleFunc(configuration.EEndpoints.GroupsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		groupListHandler.GetGroups(w, r)
	})
	http.HandleFunc(configuration.EEndpoints.TranslationEndpoint, func(w http.ResponseWriter, r *http.Request) {
		translationHandler.Translate(w, r, languageGroups)
	})
	http.HandleFunc(configuration.EEndpoints.VersionEndpoint, func(w http.ResponseWriter, r *http.Request) {
		versionHandler.GetVersion(w, r)
	})
	if configuration.Security.NeedsHTTPS {
		log.Fatal(http.ListenAndServeTLS(":8080", configuration.Security.ServerCertPath, configuration.Security.ServerKeyPath, nil))
	} else {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
