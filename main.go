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
)

func main() {
	var lReader *Config.Reader = &Config.Reader{Path: "configuration/db.json"}
	reader = lReader
	configuration, err := reader.ReadConfiguration()
	if err != nil {
		panic("Failed to read configuration")
	}
	languageGroups = configuration.Languages
	apiKey := os.Getenv("YANDEX_API_KEY")

	http.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		GroupListHandler(w, r, languageGroups)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		translationHandler = TranslationHandler{ApiKey: apiKey, Credits: configuration.Credits}
		translationHandler.Translate(w, r, languageGroups)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
