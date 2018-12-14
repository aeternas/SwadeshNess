package handlers

import (
	. "github.com/aeternas/SwadeshNess/middlewares"
)

type AnyHandler struct {
	Middlewares []Middleware
}
