package handlers

import (
	"encoding/json"
	. "github.com/aeternas/SwadeshNess/configuration"
	"log"
	"net/http"
)

type AnyVersionHandler interface {
	GetGroups(w http.ResponseWriter, r *http.Request)
}

type VersionHandler struct {
	Config *Configuration
}

func (gh *VersionHandler) GetVersion(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(gh.Config.Version)
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
