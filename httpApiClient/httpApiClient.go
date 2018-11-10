package httpApiClient

import (
	"encoding/json"
	"fmt"
	. "github.com/aeternas/SwadeshNess-packages/language"
	. "github.com/aeternas/SwadeshNess/apiClient"
	. "github.com/aeternas/SwadeshNess/configuration"
	. "github.com/aeternas/SwadeshNess/dto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type HTTPApiClient struct {
	Client      *http.Client
	Middlewares []Middleware
}

func (c *HTTPApiClient) MakeTranslationRequest(w string, conf *Configuration, sourceLang string, targetLang Language, ch chan<- YandexTranslationResult) {
	c.Middlewares = []Middleware{NewDefaultMiddleware(), NewAuthMiddleware(conf.ApiKey), NewLoggerMiddleware()}
	res := getRequest(c.Client, c.Middlewares, w, sourceLang, targetLang.Code)
	ch <- res
}

func getRequest(c *http.Client, middlewares []Middleware, w, sourceLang, targetLang string) YandexTranslationResult {
	queryString := url.QueryEscape(w)

	urlString := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?lang=%s-%s&text=%s", sourceLang, targetLang, queryString)

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		log.Println("Request initialization error: ", err)
		return getTranslationResultErrorString("Request initialization error")
	}

	for _, middleware := range middlewares {
		req = middleware.AdaptRequest(req)
	}

	resp, err := c.Do(req)

	if err != nil {
		log.Println("Request execution error: ", err)
		return getTranslationResultErrorString("Request execution error")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("I/O Read Error: ", err)
		return getTranslationResultErrorString("I/O Read Error")
	}

	var data YandexTranslationResult

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("Unmarshalling error: ", err)
		return getTranslationResultErrorString("Unmarshalling error")
	}

	defer resp.Body.Close()

	if data.Code != 200 {
		switch data.Code {
		case 401:
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
