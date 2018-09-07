package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type TranslationResult struct {
	Code int
	Text []string
}

func main() {
	apiKey := os.Getenv("YANDEX-API-KEY")
	helloHandler := func(w http.ResponseWriter, r *http.Request) {

		var reqValue string
		for k, v := range r.Form {
			if k == "tr" {
				reqValue = strings.Join(v, "")
			}
		}
		request := url.QueryEscape(reqValue)
		response := make_response(request, apiKey)
		io.WriteString(w, response)
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func make_response(w, apiKey string) string {
	client := &http.Client{Timeout: 10 * time.Second}

	urlString := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?key=%s&lang=en-ja&text=", apiKey)

	url := urlString + w

	resp, err := client.Get(url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("I/O Read Error")
	}
	var data TranslationResult
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Unmarshalling error: ", err)
	}
	if data.Code != 200 {
		fmt.Printf("Error – code is %d", data.Code)
	}
	return strings.Join(data.Text, ",")
}
