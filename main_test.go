package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestTranslationHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dev/?tr=man", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := appHandler(TranslationHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `男`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestTranslationHandlerNoAPIKey(t *testing.T) {
	if err := os.Unsetenv("YANDEX_API_KEY"); err != nil {
		t.Errorf("Failed to unset env")
	}
	req, err := http.NewRequest("GET", "/?tr=man", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := appHandler(TranslationHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v instead of error",
			status)
	}
}
