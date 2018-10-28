package configuration

import (
	. "github.com/aeternas/SwadeshNess/language"
)

const (
	API_KEY              = "YANDEX_API_KEY"
	TRANSLATION_ENDPOINT = "TRANSLATION_ENDPOINT"
	GROUP_ENDPOINT       = "GROUP_ENDPOINT"
)

type Configuration struct {
	Languages  []LanguageGroup
	ApiKey     string
	Credits    string
	EEndpoints Endpoints
}
