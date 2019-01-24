package handlers

import (
	"encoding/json"
	api "github.com/aeternas/SwadeshNess/apiClient"
	config "github.com/aeternas/SwadeshNess/configuration"
	serverMiddleware "github.com/aeternas/SwadeshNess/serverMiddlewares"
	"log"
	"net/http"
)

type VersionHandler struct {
	Config            *config.Configuration
	ServerMiddlewares []serverMiddleware.ServerMiddleware
}

func (vh *VersionHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	request := &api.Request{Data: []byte{}, Cached: false, NetRequest: r}

	for _, middleware := range vh.ServerMiddlewares {
		request = middleware.AdaptRequest(request)
	}

	bytes, err := json.Marshal(vh.Config.Version)
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

func (*VersionHandler) WriteResponse(w http.ResponseWriter, r *api.Response) {
	if _, err := w.Write(r.Data); err != nil {
		log.Println("Response output error")
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
}
