package configuration

import (
	"errors"
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"os"
	"testing"
)

var (
	osWrapper   Wrappers.AnyOsWrapper
	mockWrapper = new(MockOsWrapper)
)

func init() {
}

func TestConfigurationRead(t *testing.T) {
	file, _ := os.Open("db.json")
	mockWrapper.OpenStub = Wrappers.FileOpened{F: file, Err: errors.New("Err")}
	osWrapper = mockWrapper
	var reader *Reader = &Reader{Path: "db.json", OsWrapper: osWrapper}
	_, err := reader.ReadConfiguration()
	if err != nil && mockWrapper.OpenWasCalled == 0 && mockWrapper.OpenArgs != "db.json" {
		t.Errorf("Failed to read configuration")
	}
}

func TestLanguagesParsing(t *testing.T) {
	file, _ := os.Open("db.json")
	mockWrapper.OpenStub = Wrappers.FileOpened{F: file, Err: errors.New("Err")}
	osWrapper = mockWrapper
	var reader *Reader = &Reader{Path: "db.json", OsWrapper: osWrapper}

	config, _ := reader.ReadConfiguration()
	if len(config.Languages) < 1 {
		t.Errorf("No languages parsed from the DB")
	}
}

func TestCreditsParsing(t *testing.T) {
	file, _ := os.Open("db.json")
	mockWrapper.OpenStub = Wrappers.FileOpened{F: file, Err: errors.New("Err")}
	osWrapper = mockWrapper
	var reader *Reader = &Reader{Path: "db.json", OsWrapper: osWrapper}

	config, _ := reader.ReadConfiguration()
	if len(config.Credits) < 1 {
		t.Errorf("Invalid credits")
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
