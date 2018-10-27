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
	configuration      Config.Configuration
)

func init() {
	var lReader *Config.Reader = &Config.Reader{Path: "configuration/db.json"}
	reader = lReader
	lConfiguration, _ := reader.ReadConfiguration()
	configuration = lConfiguration
	translationHandler = &TranslationHandler{Config: &configuration}
}

func main() {
	languageGroups = configuration.Languages
	http.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		GroupListHandler(w, r, languageGroups)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		translationHandler.Translate(w, r, languageGroups)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
