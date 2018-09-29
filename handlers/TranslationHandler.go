package handlers

import (
	api "github.com/aeternas/SwadeshNess/apiClient"
	. "github.com/aeternas/SwadeshNess/dto"
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
		http.Error(w, "Please provide `group` key e.g. \"Romanic\", \"Turkic\", \"CJKV Family\"", http.StatusBadRequest)
		return
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

	ch := make(chan TranslationResult)

	for _, lang := range desiredGroup.Languages {
		go api.MakeRequest(translationRequestValue, apiKey, lang, ch)
	}

	results := []TranslationResult{}
	for range desiredGroup.Languages {
		results = append(results, <-ch)
	}

	results = getRearrangedResults(results, desiredGroup.Languages)

	translatedStrings := []string{}

	for i := range results {
		result := results[i]
		if result.Code != 200 {
			http.Error(w, result.Message, result.Code)
			return
		}

		translatedString := strings.Join(result.Text, ",")
		translatedStrings = append(translatedStrings, translatedString)
	}

	text := strings.Join(translatedStrings, "\n")

	if _, err := io.WriteString(w, text); err != nil {
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
}

func getRearrangedResults(res []TranslationResult, langs []Language) []TranslationResult {
	arrangedResults := []TranslationResult{}

	for _, desiredLang := range langs {
		for _, resultLang := range res {
			resultLangCode := strings.Split(resultLang.Lang, "-")[1]
			if desiredLang.Code == resultLangCode {
				arrangedResults = append(arrangedResults, resultLang)
				continue
			}

		}
	}
	return arrangedResults
}
