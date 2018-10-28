package main

import (
	Config "github.com/aeternas/SwadeshNess/configuration"
	. "github.com/aeternas/SwadeshNess/handlers"
	. "github.com/aeternas/SwadeshNess/language"
	"log"
	"net/http"
)

var (
	languageGroups     []LanguageGroup
	reader             Config.AnyReader
	translationHandler AnyTranslationHandler
	groupListHandler   AnyGroupListHandler
	configuration      Config.Configuration
)

func init() {
	var lReader *Config.Reader = &Config.Reader{Path: "configuration/db.json"}
	reader = lReader
	lConfiguration, _ := reader.ReadConfiguration()
	configuration = lConfiguration
	translationHandler = &TranslationHandler{Config: &configuration}
	groupListHandler = &GroupListHandler{Config: &configuration}
}

func main() {
	languageGroups = configuration.Languages
	http.HandleFunc(configuration.EEndpoints.GroupsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		groupListHandler.GetGroups(w, r)
	})
	http.HandleFunc(configuration.EEndpoints.TranslationEndpoint, func(w http.ResponseWriter, r *http.Request) {
		translationHandler.Translate(w, r, languageGroups)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
