package handlers

import (
	serverMiddleware "github.com/aeternas/SwadeshNess/serverMiddlewares"
	"net/http"
)

type AnyHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type anyHandler struct {
	Middlewares []serverMiddleware.ServerMiddleware
}

func (ah anyHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {}

func NewAnyHandler() AnyHandler {
	return &anyHandler{}
}
