package apiClient

import (
	. "github.com/aeternas/SwadeshNess/dto"
	. "github.com/aeternas/SwadeshNess/language"
)

type ApiClient interface {
	MakeTranslationRequest(w, apiKey, sourceLang string, targetLang Language, ch chan<- YandexTranslationResult)
}
