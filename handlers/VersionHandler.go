package handlers

import (
	"encoding/json"
	. "github.com/aeternas/SwadeshNess/configuration"
	serverMiddleware "github.com/aeternas/SwadeshNess/serverMiddlewares"
	"log"
	"net/http"
)

type VersionHandler struct {
	Config      *Configuration
	Middlewares []serverMiddleware.ServerMiddleware
}

func (gh *VersionHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
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
