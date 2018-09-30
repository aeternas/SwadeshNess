package main

import (
	. "github.com/aeternas/SwadeshNess/handlers"
	. "github.com/aeternas/SwadeshNess/language"
	"log"
	"net/http"
	"os"
)

var (
	turkicLanguages      = []Language{{FullName: "Tatar", Code: "tt"}, {FullName: "Bashkort", Code: "ba"}, {FullName: "Azerbaijanian", Code: "az"}, {FullName: "Turkish", Code: "tr"}}
	turkicLanguagesGroup = LanguageGroup{Name: "Turkic", Languages: turkicLanguages}

	romanianLanguages     = []Language{{FullName: "French", Code: "fr"}, {FullName: "Spanish", Code: "es"}, {FullName: "Italian", Code: "it"}, {FullName: "Romanian", Code: "ro"}}
	romanicLanguagesGroup = LanguageGroup{Name: "Romanic", Languages: romanianLanguages}

	cjkvLanguages = []Language{{FullName: "Mandarin", Code: "zh"}, {FullName: "Japanese", Code: "ja"}, {FullName: "Vietnamese", Code: "vi"}}

	cjkvLanguagesGroup = LanguageGroup{Name: "CJKV Family", Languages: cjkvLanguages}

	languageGroups = []LanguageGroup{turkicLanguagesGroup, romanicLanguagesGroup, cjkvLanguagesGroup}
)

func main() {
	apiKey := os.Getenv("YANDEX_API_KEY")

	http.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		GroupListHandler(w, r, languageGroups)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		TranslationHandler(w, r, languageGroups, apiKey)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
