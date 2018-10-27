package httpApiClient_test

import (
	. "github.com/aeternas/SwadeshNess/apiClient"
	. "github.com/aeternas/SwadeshNess/dto"
	. "github.com/aeternas/SwadeshNess/language"
	"testing"
)

type MockHTTPApiClient struct{}

func (c *MockHTTPApiClient) MakeTranslationRequest(w, apiKey, sourceLang string, targetLang Language, ch chan<- YandexTranslationResult) {
	ch <- YandexTranslationResult{Code: 200, Lang: "en-tr", Message: "", Text: []string{"adam"}}
}

func TestMakeRequest(t *testing.T) {
	apiKey := "APIKEY"

	ch := make(chan YandexTranslationResult)

	turkishLanguage := Language{FullName: "Turkish", Code: "tr"}

	var apiClient ApiClient

	httpApiClient := &MockHTTPApiClient{}

	apiClient = httpApiClient

	go apiClient.MakeTranslationRequest("man", apiKey, "en", turkishLanguage, ch)

	s := []YandexTranslationResult{}

	s = append(s, <-ch)

	if s[0].Text[0] != "adam" {
		t.Errorf("wrong translation: %v ", s)
		t.Errorf("apiKey is %v", apiKey)
	}
}
