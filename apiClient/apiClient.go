package apiClient

import (
	"encoding/json"
	"errors"
	"fmt"
	d "github.com/aeternas/SwadeshNess/dto"
	l "github.com/aeternas/SwadeshNess/language"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func MakeRequest(w string, apiKey string, l l.Language, ch chan<- string) {
	req, _, err := getRequest(w, l.Code, apiKey)
	if err != nil {
		log.Println("fail")
	}
	ch <- req
}

func getRequest(w, targetLang, apiKey string) (result string, errorCode int, resultErr error) {

	client := &http.Client{Timeout: 10 * time.Second}

	queryString := url.QueryEscape(w)

	urlString := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?key=%s&lang=en-%s&text=%s", apiKey, targetLang, queryString)

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		log.Println("Request initialization error: ", err)
		return "", 500, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Request execution error: ", err)
		return "", 500, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("I/O Read Error: ", err)
		return "", 500, err
	}

	var data d.TranslationResult

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("Unmarshalling error: ", err)
		return "", 500, err
	}

	defer resp.Body.Close()

	if data.Code != 200 {
		switch data.Code {
		case 401:
			log.Println("Invalid API Key")
			return "InvalidApiKey", data.Code, errors.New("InvalidApiKey")
		default:
			log.Printf("Error – code is %d", data.Code)
		}
	}

	return strings.Join(data.Text, ","), data.Code, nil
}
