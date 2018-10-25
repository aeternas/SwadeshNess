package handlers

import (
	"encoding/json"
	. "github.com/aeternas/SwadeshNess/language"
	"log"
	"net/http"
)

func GroupListHandler(w http.ResponseWriter, r *http.Request, groups []LanguageGroup) {
	bytes, err := json.Marshal(groups)
	if err != nil {
		log.Println("Marshalling Error")
		http.Error(w, "Marshalling error", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(bytes); err != nil {
		log.Println("Write Error")
		http.Error(w, "Response write error", http.StatusInternalServerError)
		return
	}
}
