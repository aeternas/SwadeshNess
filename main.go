package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	l "github.com/aeternas/SwadeshNess/language"
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

var (
	turkicLanguages      = []l.Language{{FullName: "Tatar", Code: "tt"}, {FullName: "Bashkort", Code: "ba"}, {FullName: "Azerbaijanian", Code: "az"}, {FullName: "Turkish", Code: "tr"}}
	turkicLanguagesGroup = l.LanguageGroup{Name: "turkic", Languages: turkicLanguages}

	romanianLanguages      = []l.Language{{FullName: "French", Code: "fr"}, {FullName: "Spanish", Code: "es"}}
	romanianLanguagesGroup = l.LanguageGroup{Name: "Romanian", Languages: romanianLanguages}
)

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
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	translationRequestValue := translationRequestValues[0]
	response, code, err := getRequest(translationRequestValue, apiKey)
	if err != nil || code != 200 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	if _, err := io.WriteString(w, response); err != nil {
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
	return nil
}

func GroupListHandler(w http.ResponseWriter, r *http.Request) error {
	var buf bytes.Buffer
	groups := []l.LanguageGroup{turkicLanguagesGroup, romanianLanguagesGroup}
	if err := json.NewEncoder(&buf).Encode(groups); err != nil {
		http.Error(w, "Encoding Error", http.StatusInternalServerError)
		return err
	}

	_, err := buf.WriteTo(w)
	if err != nil {
		http.Error(w, "Response Write Error", http.StatusInternalServerError)
		return err
	}
	return nil
}

func main() {
	http.Handle("/groups", appHandler(GroupListHandler))
	http.Handle("/", appHandler(TranslationHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRequest(w, apiKey string) (string, int, error) {

	client := &http.Client{Timeout: 10 * time.Second}

	urlString := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?key=%s&lang=en-ja&text=", apiKey)

	url := urlString + w

	req, err := http.NewRequest("GET", url, nil)
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
