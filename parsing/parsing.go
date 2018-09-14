package parsing

import (
	l "github.com/aeternas/SwadeshNess/language"
)

type TranslateRequest struct {
	Token string
	Group l.LanguageGroup
}
