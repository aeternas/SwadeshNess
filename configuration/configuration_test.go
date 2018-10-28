package configuration

import (
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"testing"
)

func TestConfigurationRead(t *testing.T) {
	var reader *Reader = &Reader{Path: "db.json", OsWrapper: new(Wrappers.OsWrapper)}
	_, err := reader.ReadConfiguration()
	if err != nil {
		t.Errorf("Failed to read configuration")
	}
}

func TestLanguagesParsing(t *testing.T) {
	var reader *Reader = &Reader{Path: "db.json", OsWrapper: new(Wrappers.OsWrapper)}
	config, _ := reader.ReadConfiguration()
	if len(config.Languages) < 1 {
		t.Errorf("No languages parsed from the DB")
	}
}

func TestCreditsParsing(t *testing.T) {
	var reader *Reader = &Reader{Path: "db.json", OsWrapper: new(Wrappers.OsWrapper)}
	config, _ := reader.ReadConfiguration()
	if len(config.Credits) < 1 {
		t.Errorf("Invalid credits")
	}
}
