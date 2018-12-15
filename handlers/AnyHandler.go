package handlers

import (
	. "github.com/aeternas/SwadeshNess/middlewares"
	"net/http"
)

type AnyHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type anyHandler struct {
	Middlewares []Middleware
}

func (ah anyHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {

}

func NewAnyHandler() AnyHandler {
	return &anyHandler{}
}
