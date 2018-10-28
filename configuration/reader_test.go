package configuration

import (
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"testing"
)

var (
	osWrapper   Wrappers.AnyOsWrapper
	mockWrapper = new(MockOsWrapper)
)

func TestReadConfiguration(t *testing.T) {
	osWrapper = mockWrapper
	var reader *Reader = &Reader{Path: "path", OsWrapper: osWrapper}
	_, err := reader.ReadConfiguration()
	if err != nil && mockWrapper.OpenWasCalled == 0 && mockWrapper.OpenArgs != "path" {
		t.Errorf("Failed to read configuration")
	}
}

type MockOsWrapper struct {
	GetEnvArgs      string
	GetEnvWasCalled int
	GetEnvStub      string

	OpenArgs      string
	OpenWasCalled int
	OpenStub      Wrappers.FileOpened
}

func (w *MockOsWrapper) GetEnv(k string) string {
	w.GetEnvWasCalled += 1
	w.GetEnvArgs = k
	return w.GetEnvStub
}

func (w *MockOsWrapper) Open(n string) Wrappers.FileOpened {
	w.OpenWasCalled += 1
	w.OpenArgs = n
	return w.OpenStub
}
