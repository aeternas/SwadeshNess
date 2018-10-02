package handlers

import (
	"errors"
	"fmt"
	. "github.com/aeternas/SwadeshNess/apiClient"
	. "github.com/aeternas/SwadeshNess/dto"
	. "github.com/aeternas/SwadeshNess/language"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func TranslationHandler(w http.ResponseWriter, r *http.Request, languageGroups []LanguageGroup, apiKey string) {
	translationRequestValues, ok := r.URL.Query()["translate"]
	if !ok || len(translationRequestValues[0]) < 1 {
		log.Println("Invalid request")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	translationRequestValue := translationRequestValues[0]

	translationRequestGroupValues, ok := r.URL.Query()["group"]
	if !ok || len(translationRequestValues[0]) < 1 {
		http.Error(w, "Please provide `group` key e.g. \"Romanic\", \"Turkic\", \"CJKV Family\"", http.StatusBadRequest)
		return
	}

	var sourceLanguage string

	sourceLanguageValues, ok := r.URL.Query()["source"]
	if !ok || len(sourceLanguageValues[0]) < 1 {
		sourceLanguage = "en"
	} else {
		sourceLanguage = sourceLanguageValues[0]
	}

	var translatedStrings []string
	for _, lang := range translationRequestGroupValues {
		res, err := getTranslation(translationRequestValue, sourceLanguage, lang, languageGroups, apiKey)
		if err != nil {
			translatedStrings = append(translatedStrings, fmt.Sprintf("Failed to process language group: %s", lang))
		} else {
			translatedStrings = append(translatedStrings, res)
		}
	}

	text := strings.Join(translatedStrings, "\n")

	if _, err := io.WriteString(w, text); err != nil {
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
}

func getTranslation(translationRequestValue, sourceLanguage, targetLanguage string, availableLanguageGroups []LanguageGroup, apiKey string) (string, error) {
	var desiredGroup LanguageGroup

	// TODO: Move to properties

	var apiClient ApiClient
	httpApiClient := &HTTPApiClient{Client: &http.Client{Timeout: 10 * time.Second}}

	apiClient = httpApiClient

	for i := range availableLanguageGroups {
		if strings.ToLower(availableLanguageGroups[i].Name) == strings.ToLower(targetLanguage) {
			desiredGroup = availableLanguageGroups[i]
			break
		}
	}

	if desiredGroup.Name == "" {
		return "", errors.New("No such language group found")
	}

	ch := make(chan TranslationResult)

	for _, lang := range desiredGroup.Languages {
		go apiClient.MakeRequest(translationRequestValue, apiKey, sourceLanguage, lang, ch)
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
			return "", errors.New(result.Message)
		}

		translatedString := strings.Join(result.Text, ",")
		translatedStrings = append(translatedStrings, translatedString)
	}

	return strings.Join(translatedStrings, "\n"), nil
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
