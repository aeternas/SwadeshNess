package main

import (
	. "github.com/aeternas/SwadeshNess-packages/language"
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
)

func init() {
	var wrapper = Wrappers.New(new(Wrappers.OsWrapper))
	var lReader *Config.Reader = &Config.Reader{Path: "configuration/db.json", OsWrapper: wrapper}
	reader = lReader
	lConfiguration, _ := reader.ReadConfiguration()
	configuration = lConfiguration
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
	log.Fatal(http.ListenAndServeTLS(":8080", "certs/certificate.chained.crt", "certs/private.key", nil))
}
