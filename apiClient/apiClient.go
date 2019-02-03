package apiClient

import (
	language "github.com/aeternas/SwadeshNess-packages/language"
	configuration "github.com/aeternas/SwadeshNess/configuration"
	dto "github.com/aeternas/SwadeshNess/dto"
)

type ApiClient interface {
	MakeTranslationRequest(w string, conf *configuration.Configuration, sourceLang string, targetLang language.Language, ch chan<- dto.YandexTranslationResult)
}
