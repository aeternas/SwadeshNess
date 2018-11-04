package wrappers

import (
	"testing"
)

func TestGetEnvFallbackSuccess(t *testing.T) {
	var mockOsWrapper = new(MockOsWrapper)
	mockOsWrapper.LookupEnvStub = "stub"
	mockOsWrapper.LookupEnvStatus = true
	actualWrapper := New(mockOsWrapper)
	res := actualWrapper.GetEnvFallback("k", "fallback")
	if mockOsWrapper.LookupEnvWasCalled != 1 {
		t.Errorf("Env wasn't lookedup: %v", mockOsWrapper.LookupEnvWasCalled)
	}
	if res != "stub" {
		t.Errorf("Env not get correctly %v", res)
	}
}

func TestGetEnvFallbackFallback(t *testing.T) {
	var mockOsWrapper = new(MockOsWrapper)
	actualWrapper := New(mockOsWrapper)
	res := actualWrapper.GetEnvFallback("k", "fallback")
	if mockOsWrapper.LookupEnvWasCalled != 1 {
		t.Errorf("Env wasn't lookedup: %v", mockOsWrapper.LookupEnvWasCalled)
	}
	if res != "fallback" {
		t.Errorf("Env not get correctly %v", res)
	}
}
