package apiClient

import (
	. "github.com/aeternas/SwadeshNess-packages/language"
	. "github.com/aeternas/SwadeshNess/dto"
)

type ApiClient interface {
	MakeTranslationRequest(w, apiKey, sourceLang string, targetLang Language, ch chan<- YandexTranslationResult)
}
