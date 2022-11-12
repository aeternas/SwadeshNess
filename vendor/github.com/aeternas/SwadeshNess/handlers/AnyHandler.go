package handlers

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	serverMiddleware "github.com/aeternas/SwadeshNess/serverMiddlewares"
	"net/http"
)

func NewAnyHandler() AnyHandler {
	return &anyHandler{}
}

type AnyHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
	WriteResponse(w http.ResponseWriter, r *apiClient.Response)
}

type anyHandler struct {
	Middlewares []serverMiddleware.ServerMiddleware
}

func (ah anyHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {}

func (ah anyHandler) WriteResponse(w http.ResponseWriter, r *apiClient.Response) {}
