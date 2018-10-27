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
	languageGroups []LanguageGroup
	reader         Config.AnyReader
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

	http.HandleFunc("/dev/groups", func(w http.ResponseWriter, r *http.Request) {
		GroupListHandler(w, r, languageGroups)
	})
	http.HandleFunc("/dev/", func(w http.ResponseWriter, r *http.Request) {
		TranslationHandler(w, r, languageGroups, apiKey)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
