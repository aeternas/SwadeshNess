package apiClient

import (
	"encoding/json"
	"fmt"
	. "github.com/aeternas/SwadeshNess/dto"
	. "github.com/aeternas/SwadeshNess/language"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func MakeRequest(w string, apiKey string, lang Language, ch chan<- TranslationResult) {
	res := getRequest(w, lang.Code, apiKey)
	ch <- res
}

func getRequest(w, targetLang, apiKey string) TranslationResult {

	client := &http.Client{Timeout: 10 * time.Second}

	queryString := url.QueryEscape(w)

	urlString := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?key=%s&lang=en-%s&text=%s", apiKey, targetLang, queryString)

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		log.Println("Request initialization error: ", err)
		return getTranslationResultErrorString("Request initialization error")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Request execution error: ", err)
		return getTranslationResultErrorString("Request execution error")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("I/O Read Error: ", err)
		return getTranslationResultErrorString("I/O Read Error")
	}

	var data TranslationResult

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

func getTranslationResultErrorString(err string) TranslationResult {
	return TranslationResult{Code: http.StatusInternalServerError, Lang: "", Message: err, Text: []string{""}}
}
