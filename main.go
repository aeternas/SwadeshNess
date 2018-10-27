package main

import (
	"encoding/json"
	. "github.com/aeternas/SwadeshNess/handlers"
	. "github.com/aeternas/SwadeshNess/language"
	"log"
	"net/http"
	"os"
)

var (
	languageGroups []LanguageGroup
	configuration  Configuration = readConfiguration()
)

type Configuration struct {
	Languages []LanguageGroup
}

func init() {
	languageGroups = configuration.Languages
}

func main() {
	readConfiguration()
	apiKey := os.Getenv("YANDEX_API_KEY")

	http.HandleFunc("/dev/groups", func(w http.ResponseWriter, r *http.Request) {
		GroupListHandler(w, r, languageGroups)
	})
	http.HandleFunc("/dev/", func(w http.ResponseWriter, r *http.Request) {
		TranslationHandler(w, r, languageGroups, apiKey)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readConfiguration() Configuration {
	file, _ := os.Open("db.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("Unable to read database")
		panic("Error reading database")
	}
	return configuration
}
