package apiClient

import (
	. "github.com/aeternas/SwadeshNess/dto"
	. "github.com/aeternas/SwadeshNess/language"
)

type ApiClient interface {
	MakeRequest(w, apiKey, sourceLang string, targetLang Language, ch chan<- YandexTranslationResult)
}
