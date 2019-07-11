package middlewaresv2

import (
	"net/http"
)

type ClientMiddlewarev2 func(http.HandlerFunc) http.HandlerFunc
