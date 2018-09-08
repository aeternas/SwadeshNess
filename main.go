package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type TranslationResult struct {
	Code int
	Text []string
}

type appHandler func(http.ResponseWriter, *http.Request) error

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func TranslationHandler(w http.ResponseWriter, r *http.Request) error {
	apiKey := os.Getenv("YANDEX_API_KEY")

	translationRequestValues, ok := r.URL.Query()["tr"]
	if !ok || len(translationRequestValues[0]) < 1 {
		log.Println("Invalid request")
	}
	translationRequestValue := translationRequestValues[0]
	response, code, err := getRequest(translationRequestValue, apiKey)
	if err != nil || code != 200 {
		fmt.Println("Error: ", err)
		return errors.New("Error is: ", err)
	}
	if _, err := io.WriteString(w, response); err != nil {
		log.Println("Response output error")
		return errors.New("Error is: ", err)
	}

	return nil
}

func main() {

	http.Handle("/", appHandler(TranslationHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRequest(w, apiKey string) (string, int, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	urlString := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?key=%s&lang=en-ja&text=", apiKey)

	url := urlString + w

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Request initialization error")
		return "", 500, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Request execution error")
		return "", 500, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("I/O Read Error")
		return "", 500, err
	}

	var data TranslationResult

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("Unmarshalling error: ", err)
		return "", 500, err
	}

	defer resp.Body.Close()

	if data.Code != 200 {
		switch data.Code {
		case 401:
			log.Println("Invalid API Key")
		default:
			log.Printf("Error – code is %d", data.Code)
		}
	}

	return strings.Join(data.Text, ","), data.Code, nil
}
