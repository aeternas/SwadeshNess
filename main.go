package main

import (
	"encoding/json"
	api "github.com/aeternas/SwadeshNess/apiClient"
	l "github.com/aeternas/SwadeshNess/language"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type appHandler func(http.ResponseWriter, *http.Request) error

var (
	turkicLanguages      = []l.Language{{FullName: "Tatar", Code: "tt"}, {FullName: "Bashkort", Code: "ba"}, {FullName: "Azerbaijanian", Code: "az"}, {FullName: "Turkish", Code: "tr"}}
	turkicLanguagesGroup = l.LanguageGroup{Name: "Turkic", Languages: turkicLanguages}

	romanianLanguages      = []l.Language{{FullName: "French", Code: "fr"}, {FullName: "Spanish", Code: "es"}}
	romanianLanguagesGroup = l.LanguageGroup{Name: "Romanian", Languages: romanianLanguages}

	languageGroups = []l.LanguageGroup{turkicLanguagesGroup, romanianLanguagesGroup}
)

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	http.Handle("/dev/groups", appHandler(GroupListHandler))
	http.Handle("/dev/", appHandler(TranslationHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TranslationHandler(w http.ResponseWriter, r *http.Request) error {
	apiKey := os.Getenv("YANDEX_API_KEY")

	translationRequestValues, ok := r.URL.Query()["tr"]
	if !ok || len(translationRequestValues[0]) < 1 {
		log.Println("Invalid request")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	translationRequestValue := translationRequestValues[0]

	var translationRequestGroupValue string

	translationRequestGroupValues, ok := r.URL.Query()["tr"]
	if !ok || len(translationRequestValues[0]) < 1 {
		translationRequestGroupValue = "Turkic"
	} else {
		translationRequestGroupValue = translationRequestGroupValues[0]
	}

	var desiredGroup l.LanguageGroup

	for i := range languageGroups {
		if languageGroups[i].Name == translationRequestGroupValue {
			desiredGroup = languageGroups[i]
			break
		}
	}

	ch := make(chan string)

	for _, lang := range desiredGroup.Languages {
		go api.MakeRequest(translationRequestValue, apiKey, lang, ch)
	}

	s := []string{}
	for range turkicLanguages {
		s = append(s, <-ch)
	}

	response := strings.Join(s, "\n")

	if _, err := io.WriteString(w, response); err != nil {
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
	return nil
}

func GroupListHandler(w http.ResponseWriter, r *http.Request) error {
	groups := []l.LanguageGroup{turkicLanguagesGroup, romanianLanguagesGroup}

	bytes, err := json.Marshal(groups)
	if err != nil {
		http.Error(w, "Marshalling error", http.StatusInternalServerError)
		return err
	}

	if _, err := w.Write(bytes); err != nil {
		http.Error(w, "Response write error", http.StatusInternalServerError)
		return err
	}

	return nil
}
