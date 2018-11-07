package parsing

import (
	l "github.com/aeternas/SwadeshNess-packages/language"
)

type TranslateRequest struct {
	Token string
	Group l.LanguageGroup
}
