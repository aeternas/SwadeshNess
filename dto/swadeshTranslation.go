package translation

type SwadeshTranslation struct {
	Results []GroupTranslation
}

type LanguageTranslation struct {
	Name        string
	Translation string
}

type GroupTranslation struct {
	Name    string
	Results []LanguageTranslation
}
