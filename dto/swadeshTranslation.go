package translation

type SwadeshTranslation struct {
	Results []GroupTranslation `json:"results"`
}

type LanguageTranslation struct {
	Name        string `json:"name"`
	Translation string `json:"translation"`
}

type GroupTranslation struct {
	Name    string                `json:"name"`
	Results []LanguageTranslation `json:"results"`
}
