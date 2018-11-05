package configuration

import (
	. "github.com/aeternas/SwadeshNess/language"
)

const (
	API_KEY              = "YANDEX_API_KEY"
	TRANSLATION_ENDPOINT = "TRANSLATION_ENDPOINT"
	GROUP_ENDPOINT       = "GROUP_ENDPOINT"
	VERSION_ENDPOINT     = "VERSION_ENDPOINT"
	VERSION              = "VERSION"
)

type Configuration struct {
	Languages  []LanguageGroup
	ApiKey     string
	Credits    string
	Version    string
	EEndpoints Endpoints
}
