package httpApiClient

import (
	"encoding/json"
	"fmt"
	language "github.com/aeternas/SwadeshNess-packages/language"
	ApiClient "github.com/aeternas/SwadeshNess/apiClient"
	middlewares "github.com/aeternas/SwadeshNess/clientMiddlewares"
	configuration "github.com/aeternas/SwadeshNess/configuration"
	. "github.com/aeternas/SwadeshNess/dto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type HTTPApiClient struct {
	Client      *http.Client
	Middlewares []middlewares.ClientMiddleware
}

func (c *HTTPApiClient) MakeTranslationRequest(w string, conf *configuration.Configuration, sourceLang string, targetLang language.Language, ch chan<- YandexTranslationResult) {
	res := c.getRequest(c.Middlewares, w, sourceLang, targetLang.Code)
	ch <- res
}

func (c *HTTPApiClient) getRequest(middlewares []middlewares.ClientMiddleware, w, sourceLang, targetLang string) YandexTranslationResult {
	queryString := url.QueryEscape(w)

	urlString := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?lang=%s-%s&text=%s", sourceLang, targetLang, queryString)

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Println("Request initialization error: ", err)
		return getTranslationResultErrorString("Request initialization error")
	}

	request := &ApiClient.Request{Data: []byte{}, Cached: false, NetRequest: req}

	for _, middleware := range middlewares {
		request = middleware.AdaptRequest(request)
	}

	response := &ApiClient.Response{Data: []byte{}, NetResponse: nil, Request: request}

	if request.Cached {
		response = c.adaptResponse(response)
		return c.getTranslationData(response)
	}

	resp, err := c.Client.Do(request.NetRequest)

	if err != nil {
		log.Println("Request execution error: ", err)
		return getTranslationResultErrorString("Request execution error")
	}

	response.NetResponse = resp

	body, err := ioutil.ReadAll(response.NetResponse.Body)

	if err != nil {
		log.Println("I/O Read Error: ", err)
		return getTranslationResultErrorString("I/O Read Error")
	}

	response.Data = body

	response = c.adaptResponse(response)

	defer response.NetResponse.Body.Close()

	return c.getTranslationData(response)
}

func (c *HTTPApiClient) adaptResponse(r *ApiClient.Response) *ApiClient.Response {
	adaptedResponse := r
	for _, middleware := range c.Middlewares {
		adaptedResponse = middleware.AdaptResponse(adaptedResponse)
	}

	return adaptedResponse
}

func (c *HTTPApiClient) getTranslationData(r *ApiClient.Response) YandexTranslationResult {
	var data YandexTranslationResult

	if err := json.Unmarshal(r.Data, &data); err != nil {
		log.Println("Unmarshalling error: ", err)
		return getTranslationResultErrorString("Unmarshalling error")
	}

	if data.Code != http.StatusOK {
		switch data.Code {
		case http.StatusUnauthorized:
			log.Println("Invalid API Key")
			return getTranslationResultErrorString("Invalid API Key")
		default:
			log.Printf("Error – code is %d", data.Code)
		}
	}

	return data
}

func getTranslationResultErrorString(err string) YandexTranslationResult {
	return YandexTranslationResult{Code: http.StatusInternalServerError, Lang: "", Message: err, Text: []string{""}}
}
