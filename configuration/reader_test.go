package configuration

import (
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"reflect"
	"testing"
)

var (
	expectedGetEnvArgs         = []string{"YANDEX_API_KEY"}
	expectedGetEnvFallbackArgs = []string{"TRANSLATION_ENDPOINT", "/v1/", "GROUP_ENDPOINT", "/v1/groups", "VERSION_ENDPOINT", "/v1/version", "VERSION", "0", "SERVER_KEY", "certs/server.key", "SERVER_CERT", "certs/server.crt", "REDIS_ADDRESS", "localhost"}
	osWrapper                  Wrappers.AnyOsWrapper
	mockWrapper                = new(Wrappers.MockOsWrapper)
)

func TestReadConfiguration(t *testing.T) {
	osWrapper = mockWrapper
	var reader *Reader = &Reader{Path: "path", OsWrapper: osWrapper}
	_, err := reader.ReadConfiguration()
	if mockWrapper.OpenWasCalled != 1 {
		t.Errorf("Open was not called %v", mockWrapper.OpenWasCalled)
	}
	if mockWrapper.OpenArgs != "path" {
		t.Errorf("Open path is not valid %v", mockWrapper.OpenArgs)
	}
	if mockWrapper.GetEnvWasCalled != 1 {
		t.Errorf("GetEnv was not called valid amount of times: %v instead of 1", mockWrapper.GetEnvWasCalled)
	}
	if !reflect.DeepEqual(mockWrapper.GetEnvArgs, expectedGetEnvArgs) {
		t.Errorf("GetEnv total args are not equal to expected, %v", mockWrapper.GetEnvArgs)
	}
	if mockWrapper.GetEnvFallbackWasCalled != 7 {
		t.Errorf("GetEnvFallback was not called valid amount of times: %v instead of 7", mockWrapper.GetEnvFallbackWasCalled)
	}
	if !reflect.DeepEqual(mockWrapper.GetEnvFallbackArgs, expectedGetEnvFallbackArgs) {
		t.Errorf("GetEnvFallback args are %v", mockWrapper.GetEnvFallbackArgs)
	}
	if err == nil {
		t.Errorf("No error occured, test logic failure")
	}
}
