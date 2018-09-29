package handlers

import (
	api "github.com/aeternas/SwadeshNess/apiClient"
	. "github.com/aeternas/SwadeshNess/language"
	"io"
	"log"
	"net/http"
	"strings"
)

func TranslationHandler(w http.ResponseWriter, r *http.Request, languageGroups []LanguageGroup, apiKey string) {
	translationRequestValues, ok := r.URL.Query()["translate"]
	if !ok || len(translationRequestValues[0]) < 1 {
		log.Println("Invalid request")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	translationRequestValue := translationRequestValues[0]

	var translationRequestGroupValue string

	translationRequestGroupValues, ok := r.URL.Query()["group"]
	if !ok || len(translationRequestValues[0]) < 1 {
		translationRequestGroupValue = "Turkic"
	} else {
		translationRequestGroupValue = translationRequestGroupValues[0]
	}

	var desiredGroup LanguageGroup

	for i := range languageGroups {
		if strings.ToLower(languageGroups[i].Name) == strings.ToLower(translationRequestGroupValue) {
			desiredGroup = languageGroups[i]
			break
		}
	}

	if desiredGroup.Name == "" {
		http.Error(w, "No such language group found", http.StatusBadRequest)
		return
	}

	ch := make(chan string)

	for _, lang := range desiredGroup.Languages {
		go api.MakeRequest(translationRequestValue, apiKey, lang, ch)
	}

	s := []string{}
	for range desiredGroup.Languages {
		s = append(s, <-ch)
	}

	response := strings.Join(s, "\n")

	if _, err := io.WriteString(w, response); err != nil {
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
}
