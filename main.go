package main

import (
	"encoding/json"
	"fmt"
	l "github.com/aeternas/SwadeshNess/language"
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

func MakeRequest(w string, apiKey string, l l.Language, ch chan<- string) {
	req, _, _ := getRequest(w, l.Code, apiKey)
	ch <- req
}

func TranslationHandler(w http.ResponseWriter, r *http.Request) error {
	apiKey := os.Getenv("YANDEX_API_KEY")

	translationRequestValues, ok := r.URL.Query()["tr"]
	if !ok || len(translationRequestValues[0]) < 1 {
		log.Println("Invalid request")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	translationRequestValue := translationRequestValues[0]

	ch := make(chan string)

	for _, lang := range turkicLanguages {
		go MakeRequest(translationRequestValue, apiKey, lang, ch)
	}

	s := []string{}
	for range turkicLanguages {
		s = append(s, <-ch)
	}

	response := strings.Join(s, "\n")

	if _, err := io.WriteString(w, response); err != nil {
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
	return nil
}

func GroupListHandler(w http.ResponseWriter, r *http.Request) error {
	groups := []l.LanguageGroup{turkicLanguagesGroup, romanianLanguagesGroup}

	bytes, err := json.Marshal(groups)
	if err != nil {
		http.Error(w, "Marshalling error", http.StatusInternalServerError)
		return err
	}

	if _, err := w.Write(bytes); err != nil {
		http.Error(w, "Response write error", http.StatusInternalServerError)
		return err
	}

	return nil
}

func main() {
	http.Handle("/dev/groups", appHandler(GroupListHandler))
	http.Handle("/dev/", appHandler(TranslationHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
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
