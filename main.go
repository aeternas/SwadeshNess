package main

import (
	"encoding/json"
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

func TranslationHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("YANDEX_API_KEY")

	translationRequestValues, ok := r.URL.Query()["tr"]
	if !ok || len(translationRequestValues[0]) < 1 {
		log.Println("Invalid request")
	}
	translationRequestValue := translationRequestValues[0]
	response := getRequest(translationRequestValue, apiKey)
	if _, err := io.WriteString(w, response); err != nil {
		log.Println("Response output error")
	}
}

func main() {

	http.HandleFunc("/", TranslationHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRequest(w, apiKey string) string {
	client := &http.Client{Timeout: 10 * time.Second}

	urlString := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?key=%s&lang=en-ja&text=", apiKey)

	url := urlString + w

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Request initialization error")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Request execution error")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("I/O Read Error")
	}

	var data TranslationResult

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("Unmarshalling error: ", err)
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

	return strings.Join(data.Text, ",")
}
