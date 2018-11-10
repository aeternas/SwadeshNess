package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/aeternas/SwadeshNess-packages/dto"
	. "github.com/aeternas/SwadeshNess-packages/language"
	. "github.com/aeternas/SwadeshNess/apiClient"
	. "github.com/aeternas/SwadeshNess/configuration"
	. "github.com/aeternas/SwadeshNess/dto"
	. "github.com/aeternas/SwadeshNess/httpApiClient"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	apiClient ApiClient
)

func init() {
	httpApiClient := &HTTPApiClient{Client: &http.Client{Timeout: 10 * time.Second}}
	apiClient = httpApiClient
}

type AnyTranslationHandler interface {
	Translate(w http.ResponseWriter, r *http.Request, languageGroups []LanguageGroup)
}

type TranslationHandler struct {
	Config *Configuration
}

func (th *TranslationHandler) Translate(w http.ResponseWriter, r *http.Request, languageGroups []LanguageGroup) {
	translationRequestValues, ok := r.URL.Query()["translate"]
	if !ok || len(translationRequestValues[0]) < 1 {
		log.Printf("Invalid Request: %s", r.URL.String())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	translationRequestValue := translationRequestValues[0]

	translationRequestGroupValues, ok := r.URL.Query()["group"]
	if !ok || len(translationRequestValues[0]) < 1 {
		log.Println("No Group Requested")
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

	groups := []GroupTranslation{}
	for _, lang := range translationRequestGroupValues {
		res, err := getTranslation(translationRequestValue, sourceLanguage, lang, th.Config)
		if err != nil {
			log.Printf("Failed to process language group: %s", lang)
			http.Error(w, fmt.Sprintf("Failed to process language group: %s", lang), http.StatusInternalServerError)
			return
		} else {
			groups = append(groups, res.Results[0])
		}
	}

	swadeshTranslation := SwadeshTranslation{Results: groups, Credits: th.Config.Credits}

	bytes, err := json.Marshal(swadeshTranslation)
	if err != nil {
		log.Println("Marshalling Error")
		http.Error(w, "Failed to marshall translation result response", http.StatusInternalServerError)
	}
	if _, err := w.Write(bytes); err != nil {
		log.Println("Response output error")
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
}

func getTranslation(translationRequestValue, sourceLanguage, targetLanguage string, conf *Configuration) (SwadeshTranslation, error) {
	var desiredGroup LanguageGroup

	for i := range conf.Languages {
		if strings.ToLower(conf.Languages[i].Name) == strings.ToLower(targetLanguage) {
			desiredGroup = conf.Languages[i]
			break
		}
	}

	if desiredGroup.Name == "" {
		return SwadeshTranslation{Results: []GroupTranslation{}}, errors.New("No such language group found")
	}

	ch := make(chan YandexTranslationResult)

	for _, lang := range desiredGroup.Languages {
		go apiClient.MakeTranslationRequest(translationRequestValue, conf, sourceLanguage, lang, ch)
	}

	results := []YandexTranslationResult{}
	for range desiredGroup.Languages {
		results = append(results, <-ch)
	}

	swadeshResults := translateToSwadeshTranslation(results, desiredGroup, conf.Credits)

	return swadeshResults, nil
}

func translateToSwadeshTranslation(res []YandexTranslationResult, desiredGroup LanguageGroup, credits string) SwadeshTranslation {

	languageTranslationResult := []LanguageTranslation{}

	for _, desiredLang := range desiredGroup.Languages {
		for _, yandexResult := range res {
			resultLangCode := strings.Split(yandexResult.Lang, "-")[1]
			if desiredLang.Code == resultLangCode && yandexResult.Code == 200 {
				languageTranslationResult = append(languageTranslationResult, LanguageTranslation{Name: desiredLang.FullName, Translation: strings.Join(yandexResult.Text, ",")})
				continue
			}

		}
	}

	groupTranslationResult := []GroupTranslation{{Name: desiredGroup.Name, Results: languageTranslationResult}}
	swadeshTranslation := SwadeshTranslation{Results: groupTranslationResult, Credits: credits}

	return swadeshTranslation
}
