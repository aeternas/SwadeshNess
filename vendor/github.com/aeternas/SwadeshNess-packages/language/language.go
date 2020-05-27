package language

type Language struct {
	FullName string `json:"fullName"`
	Code     string `json:"code"`
}

type LanguageGroup struct {
	Name      string     `json:"name"`
	Languages []Language `json:"languages"`
}
