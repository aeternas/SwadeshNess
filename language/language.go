package language

type Language struct {
	FullName string
	Code     string
}

type LanguageGroup struct {
	Name      string
	Languages []Language
}
