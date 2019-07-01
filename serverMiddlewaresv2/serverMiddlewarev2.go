package middlewaresv2

import (
	"net/http"
)

type Middlewarev2 func(http.HandlerFunc) http.HandlerFunc
