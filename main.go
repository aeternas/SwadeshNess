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

func main() {
	apiKey := os.Getenv("YANDEX_API_KEY")
	helloHandler := func(w http.ResponseWriter, r *http.Request) {

		translationRequestValues, ok := r.URL.Query()["tr"]
		if !ok || len(translationRequestValues[0]) < 1 {
			log.Println("Invalid request")
		}
		translationRequestValue := translationRequestValues[0]
		response := make_response(translationRequestValue, apiKey)
		if _, err := io.WriteString(w, response); err != nil {
			fmt.Println("Response output error")
		}
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func make_response(w, apiKey string) string {
	client := &http.Client{Timeout: 10 * time.Second}

	urlString := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?key=%s&lang=en-ja&text=", apiKey)

	url := urlString + w

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Request initialization error")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Request execution error")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("I/O Read Error")
	}

	var data TranslationResult

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Unmarshalling error: ", err)
	}

	defer resp.Body.Close()

	if data.Code != 200 {
		switch data.Code {
		case 401:
			fmt.Println("Invalid API Key")
		default:
			fmt.Printf("Error – code is %d", data.Code, data)
		}
	}

	return strings.Join(data.Text, ",")
}
