package handlers

import (
	serverMiddleware "github.com/aeternas/SwadeshNess/serverMiddlewares"
	"net/http"
)

func NewAnyHandler() AnyHandler {
	return &anyHandler{}
}

type AnyHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type anyHandler struct {
	Middlewares []serverMiddleware.ServerMiddleware
}

func (ah anyHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {}
