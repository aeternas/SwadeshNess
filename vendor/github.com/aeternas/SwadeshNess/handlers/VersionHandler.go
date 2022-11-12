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

	resp := &api.Response{Data: bytes, NetResponse: nil, Request: request}

	adaptedResponse := vh.adaptResponse(resp)

	vh.WriteResponse(w, adaptedResponse)
}

func (vh *VersionHandler) adaptResponse(r *api.Response) *api.Response {
	adaptedResponse := r
	for _, middleware := range vh.ServerMiddlewares {
		adaptedResponse = middleware.AdaptResponse(adaptedResponse)
	}

	return adaptedResponse
}

func (*VersionHandler) WriteResponse(w http.ResponseWriter, r *api.Response) {
	if _, err := w.Write(r.Data); err != nil {
		log.Println("Response output error")
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
}
