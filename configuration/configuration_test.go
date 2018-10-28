package configuration

import (
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"testing"
)

var (
	osWrapper Wrappers.AnyOsWrapper
)

func init() {
	osWrapper = new(MockOsWrapper)
}

func TestConfigurationRead(t *testing.T) {
	var reader *Reader = &Reader{Path: "db.json", OsWrapper: osWrapper}
	_, err := reader.ReadConfiguration()
	if err != nil {
		t.Errorf("Failed to read file")
	}
}

func TestLanguagesParsing(t *testing.T) {
	var reader *Reader = &Reader{Path: "db.json", OsWrapper: osWrapper}

	config, _ := reader.ReadConfiguration()
	if len(config.Languages) < 1 {
		t.Errorf("No languages parsed from the DB")
	}
}

func TestCreditsParsing(t *testing.T) {
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
}

func (w *MockOsWrapper) GetEnv(k string) string {
	w.GetEnvWasCalled += 1
	w.GetEnvArgs = k
	return w.GetEnvStub
}
