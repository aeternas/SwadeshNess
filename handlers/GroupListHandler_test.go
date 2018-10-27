package handlers

import (
	. "github.com/aeternas/SwadeshNess/language"
	"testing"
)

func TestGroupListHandler(t *testing.T) {
	t.SkipNow()
	gps := []LanguageGroup{LanguageGroup{Name: "Group:", Languages: []Language{{FullName: "Language", Code: "Lang"}}}}
	if len(gps) > 0 {
		t.Errorf("Err")
	}
}
