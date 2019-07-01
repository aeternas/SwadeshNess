package main

type Middlewarev2 func(http.HandlerFunc) http.HandlerFunc
