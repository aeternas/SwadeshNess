package handlers

import (
	"encoding/json"
	. "github.com/aeternas/SwadeshNess/apiClient"
	. "github.com/aeternas/SwadeshNess/configuration"
	serverMiddleware "github.com/aeternas/SwadeshNess/serverMiddlewares"
	"log"
	"net/http"
)

type GroupListHandler struct {
	Config            *Configuration
	ServerMiddlewares []serverMiddleware.ServerMiddleware
}

func (gh *GroupListHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	request := &Request{Data: []byte{}, Cached: false, NetRequest: r}

	for _, middleware := range gh.ServerMiddlewares {
		request = middleware.AdaptRequest(request)
	}

	bytes, err := json.Marshal(gh.Config.Languages)
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
