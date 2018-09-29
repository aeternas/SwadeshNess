package handlers

import (
	"encoding/json"
	. "github.com/aeternas/SwadeshNess/language"
	"net/http"
)

func GroupListHandler(w http.ResponseWriter, r *http.Request, groups []LanguageGroup) {
	bytes, err := json.Marshal(groups)
	if err != nil {
		http.Error(w, "Marshalling error", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(bytes); err != nil {
		http.Error(w, "Response write error", http.StatusInternalServerError)
		return
	}
}
