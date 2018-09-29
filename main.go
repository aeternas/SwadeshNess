package main

import (
	. "github.com/aeternas/SwadeshNess/handlers"
	l "github.com/aeternas/SwadeshNess/language"
	"log"
	"net/http"
	"os"
)

var (
	turkicLanguages      = []l.Language{{FullName: "Tatar", Code: "tt"}, {FullName: "Bashkort", Code: "ba"}, {FullName: "Azerbaijanian", Code: "az"}, {FullName: "Turkish", Code: "tr"}}
	turkicLanguagesGroup = l.LanguageGroup{Name: "Turkic", Languages: turkicLanguages}

	romanianLanguages      = []l.Language{{FullName: "French", Code: "fr"}, {FullName: "Spanish", Code: "es"}, {FullName: "Italian", Code: "it"}, {FullName: "Romanian", Code: "ro"}}
	romanianLanguagesGroup = l.LanguageGroup{Name: "Romanian", Languages: romanianLanguages}

	cjkvLanguages = []l.Language{{FullName: "Mandarin", Code: "zh"}, {FullName: "Japanese", Code: "ja"}, {FullName: "Vietnamese", Code: "vi"}}

	cjkvLanguagesGroup = l.LanguageGroup{Name: "CJKV Family", Languages: cjkvLanguages}

	languageGroups = []l.LanguageGroup{turkicLanguagesGroup, romanianLanguagesGroup, cjkvLanguagesGroup}
)

func main() {
	apiKey := os.Getenv("YANDEX_API_KEY")

	http.HandleFunc("/dev/groups", func(w http.ResponseWriter, r *http.Request) {
		GroupListHandler(w, r, languageGroups)
	})
	http.HandleFunc("/dev/", func(w http.ResponseWriter, r *http.Request) {
		TranslationHandler(w, r, languageGroups, apiKey)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
