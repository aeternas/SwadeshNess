package apiClient

import (
	. "github.com/aeternas/SwadeshNess/dto"
	l "github.com/aeternas/SwadeshNess/language"
	"os"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	apiKey := os.Getenv("YANDEX_API_KEY")

	ch := make(chan TranslationResult)

	turkishLanguage := l.Language{FullName: "Turkish", Code: "tr"}

	go MakeRequest("man", apiKey, turkishLanguage, ch)

	s := []TranslationResult{}

	s = append(s, <-ch)

	if s[0].Text[0] != "adam" {
		t.Errorf("wrong translation: %v ", s)
		t.Errorf("apiKey is %v", apiKey)
	}
}
