package main

import (
	. "github.com/aeternas/SwadeshNess/configuration"
	. "github.com/aeternas/SwadeshNess/handlers"
	. "github.com/aeternas/SwadeshNess/language"
	"log"
	"net/http"
	"os"
)

var (
	languageGroups []LanguageGroup
)

func main() {
	configuration, err := ReadConfiguration()
	if err != nil {
		panic("Failed to read configuration")
	}
	languageGroups = configuration.Languages
	apiKey := os.Getenv("YANDEX_API_KEY")

	http.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		GroupListHandler(w, r, languageGroups)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		TranslationHandler(w, r, languageGroups, apiKey)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
