package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	dto "github.com/aeternas/SwadeshNess-packages/dto"
	language "github.com/aeternas/SwadeshNess-packages/language"
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	configuration "github.com/aeternas/SwadeshNess/configuration"
	yandexDTO "github.com/aeternas/SwadeshNess/dto"
	serverMiddleware "github.com/aeternas/SwadeshNess/serverMiddlewares"
	"log"
	"net/http"
	"strings"
)

type TranslationHandler struct {
	Config            *configuration.Configuration
	ServerMiddlewares []serverMiddleware.ServerMiddleware
	ApiClient         apiClient.ApiClient
}

func (th *TranslationHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	translationRequestValues, ok := r.URL.Query()["translate"]
	if !ok || len(translationRequestValues[0]) < 1 {
		log.Println("Invalid Request: ", r.URL.String())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	request := &apiClient.Request{Data: []byte{}, Cached: false, NetRequest: r}

	for _, middleware := range th.ServerMiddlewares {
		request = middleware.AdaptRequest(request)
	}

	if request.Cached {
		response := &apiClient.Response{Data: []byte{}, NetResponse: nil, Request: request}
		response = th.adaptResponse(response)
		th.WriteResponse(w, response)
		return
	}

	translationRequestValue := translationRequestValues[0]

	translationRequestGroupValues, ok := request.NetRequest.URL.Query()["group"]
	if !ok || len(translationRequestValues[0]) < 1 {
		log.Println("No Group Requested")
		http.Error(w, "Please provide `group` key e.g. \"Romanic\", \"Turkic\", \"CJKV Family\"", http.StatusBadRequest)
		return
	}

	var sourceLanguage string

	sourceLanguageValues, ok := request.NetRequest.URL.Query()["source"]
	if !ok || len(sourceLanguageValues[0]) < 1 {
		sourceLanguage = "en"
	} else {
		sourceLanguage = sourceLanguageValues[0]
	}

	groups := []dto.GroupTranslation{}
	for _, lang := range translationRequestGroupValues {
		res, err := th.getTranslation(translationRequestValue, sourceLanguage, lang, th.Config)
		if err != nil {
			log.Println("Failed to process language group: ", lang)
			http.Error(w, fmt.Sprintln("Failed to process language group: ", lang), http.StatusInternalServerError)
			return
		} else {
			groups = append(groups, res.Results[0])
		}
	}

	swadeshTranslation := dto.SwadeshTranslation{Results: groups, Credits: th.Config.Credits}

	bytes, err := json.Marshal(swadeshTranslation)
	if err != nil {
		log.Println("Marshalling Error")
		http.Error(w, "Failed to marshall translation result response", http.StatusInternalServerError)
	}

	resp := &apiClient.Response{Data: bytes, NetResponse: nil, Request: request}

	adaptedResponse := th.adaptResponse(resp)

	th.WriteResponse(w, adaptedResponse)
}

func (th *TranslationHandler) adaptResponse(r *apiClient.Response) *apiClient.Response {
	adaptedResponse := r
	for _, middleware := range th.ServerMiddlewares {
		adaptedResponse = middleware.AdaptResponse(adaptedResponse)
	}

	return adaptedResponse
}

func (*TranslationHandler) WriteResponse(w http.ResponseWriter, r *apiClient.Response) {
	if _, err := w.Write(r.Data); err != nil {
		log.Println("Response output error")
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
}

func (th *TranslationHandler) getTranslation(translationRequestValue, sourceLanguage, targetLanguage string, conf *configuration.Configuration) (dto.SwadeshTranslation, error) {
	var desiredGroup language.LanguageGroup

	for i := range conf.Languages {
		if strings.ToLower(conf.Languages[i].Name) == strings.ToLower(targetLanguage) {
			desiredGroup = conf.Languages[i]
			break
		}
	}

	if desiredGroup.Name == "" {
		return dto.SwadeshTranslation{Results: []dto.GroupTranslation{}}, errors.New("No such language group found")
	}

	ch := make(chan yandexDTO.YandexTranslationResult)

	for _, lang := range desiredGroup.Languages {
		go th.ApiClient.MakeTranslationRequest(translationRequestValue, conf, sourceLanguage, lang, ch)
	}

	results := []yandexDTO.YandexTranslationResult{}
	for range desiredGroup.Languages {
		results = append(results, <-ch)
	}

	swadeshResults := translateToSwadeshTranslation(results, desiredGroup, conf.Credits)

	return swadeshResults, nil
}

func translateToSwadeshTranslation(res []yandexDTO.YandexTranslationResult, desiredGroup language.LanguageGroup, credits string) dto.SwadeshTranslation {

	languageTranslationResult := []dto.LanguageTranslation{}

	for _, desiredLang := range desiredGroup.Languages {
		for _, yandexResult := range res; yandexResult.Code == 200 {
			log.Println("Yandex result: ", yandexResult)
			resultLangCodePair := strings.Split(yandexResult.Lang, "-")
			log.Println("Result lang code pair: ", resultLangCodePair)
			resultLangCode := resultLangCodePair[1]
			if desiredLang.Code == resultLangCode && yandexResult.Code == http.StatusOK {
				languageTranslationResult = append(languageTranslationResult, dto.LanguageTranslation{Name: desiredLang.FullName, Translation: strings.Join(yandexResult.Text, ",")})
				continue
			}
		}
	}

	groupTranslationResult := []dto.GroupTranslation{{Name: desiredGroup.Name, Results: languageTranslationResult}}
	swadeshTranslation := dto.SwadeshTranslation{Results: groupTranslationResult, Credits: credits}

	return swadeshTranslation
}
