package handlers

import (
	"encoding/json"
	api "github.com/aeternas/SwadeshNess/apiClient"
	config "github.com/aeternas/SwadeshNess/configuration"
	serverMiddleware "github.com/aeternas/SwadeshNess/serverMiddlewares"
	"log"
	"net/http"
)

type GroupListHandler struct {
	Config            *config.Configuration
	ServerMiddlewares []serverMiddleware.ServerMiddleware
}

func (gh *GroupListHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	request := &api.Request{Data: []byte{}, Cached: false, NetRequest: r}

	for _, middleware := range gh.ServerMiddlewares {
		request = middleware.AdaptRequest(request)
	}

	bytes, err := json.Marshal(gh.Config.Languages)
	if err != nil {
		log.Println("Marshalling Error: ", err)
		http.Error(w, "Marshalling error", http.StatusInternalServerError)
		return
	}

	resp := &api.Response{Data: bytes, NetResponse: nil, Request: request}

	adaptedResponse := gh.adaptResponse(resp)

	gh.WriteResponse(w, adaptedResponse)
}

func (gh *GroupListHandler) adaptResponse(r *api.Response) *api.Response {
	adaptedResponse := r
	for _, middleware := range gh.ServerMiddlewares {
		adaptedResponse = middleware.AdaptResponse(adaptedResponse)
	}

	return adaptedResponse
}

func (*GroupListHandler) WriteResponse(w http.ResponseWriter, r *api.Response) {
	if _, err := w.Write(r.Data); err != nil {
		log.Println("Response output error: ", err)
		http.Error(w, "Response output error", http.StatusInternalServerError)
	}
}
