package configuration

import (
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"reflect"
	"testing"
)

var (
	expectedGetEnvArgs = []string{"YANDEX_API_KEY", "TRANSLATION_ENDPOINT", "GROUP_ENDPOINT"}
	osWrapper          Wrappers.AnyOsWrapper
	mockWrapper        = new(MockOsWrapper)
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
	if mockWrapper.GetEnvWasCalled != 3 {
		t.Errorf("GetEnv was not called valid amount of times: %v instead of 3", mockWrapper.GetEnvWasCalled)
	}
	if reflect.DeepEqual(mockWrapper.GetEnvArgs, expectedGetEnvArgs) == false {
		t.Errorf("GetEnv total args are not equal to expected, %v", mockWrapper.GetEnvArgs)
	}
	if err == nil {
		t.Errorf("No error occured, test logic failure")
	}
}

type MockOsWrapper struct {
	GetEnvArgs      []string
	GetEnvWasCalled int
	GetEnvStub      string

	OpenArgs      string
	OpenWasCalled int
	OpenStub      Wrappers.FileOpened
}

func (w *MockOsWrapper) GetEnv(k string) string {
	w.GetEnvWasCalled += 1
	w.GetEnvArgs = append(w.GetEnvArgs, k)
	return w.GetEnvStub
}

func (w *MockOsWrapper) Open(n string) Wrappers.FileOpened {
	w.OpenWasCalled += 1
	w.OpenArgs = n
	return w.OpenStub
}
