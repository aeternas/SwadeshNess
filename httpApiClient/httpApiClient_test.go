package httpApiClient

import (
	. "github.com/aeternas/SwadeshNess/apiClient"
	. "github.com/aeternas/SwadeshNess/dto"
	l "github.com/aeternas/SwadeshNess/language"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMakeRequest(t *testing.T) {
	apiKey := os.Getenv("YANDEX_API_KEY")

	ch := make(chan TranslationResult)

	turkishLanguage := l.Language{FullName: "Turkish", Code: "tr"}

	var apiClient ApiClient

	httpApiClient := &HTTPApiClient{Client: &http.Client{Timeout: 10 * time.Second}}

	apiClient = httpApiClient

	go apiClient.MakeRequest("man", apiKey, "en", turkishLanguage, ch)

	s := []TranslationResult{}

	s = append(s, <-ch)

	if s[0].Text[0] != "adam" {
		t.Errorf("wrong translation: %v ", s)
		t.Errorf("apiKey is %v", apiKey)
	}
}
