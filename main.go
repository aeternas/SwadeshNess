package main

import (
	Config "github.com/aeternas/SwadeshNess/configuration"
	. "github.com/aeternas/SwadeshNess/handlers"
	. "github.com/aeternas/SwadeshNess/language"
	"log"
	"net/http"
	"os"
)

var (
	languageGroups     []LanguageGroup
	reader             Config.AnyReader
	translationHandler AnyTranslationHandler
	configuration      Config.Configuration
)

func init() {
	var lReader *Config.Reader = &Config.Reader{Path: "configuration/db.json"}
	apiKey := os.Getenv("YANDEX_API_KEY")
	reader = lReader
	lConfiguration, _ := reader.ReadConfiguration()
	configuration = lConfiguration
	translationHandler = TranslationHandler{ApiKey: apiKey, Credits: configuration.Credits}
}

func main() {
	languageGroups = configuration.Languages
	http.HandleFunc("/dev/groups", func(w http.ResponseWriter, r *http.Request) {
		GroupListHandler(w, r, languageGroups)
	})
	http.HandleFunc("/dev/", func(w http.ResponseWriter, r *http.Request) {
		translationHandler.Translate(w, r, languageGroups)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
