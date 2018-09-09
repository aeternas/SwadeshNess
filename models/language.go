package main

import "fmt"

type Language struct {
	FullName     string
	Abbreviation string
}

type LanguageGroup struct {
	Languages []Language
}
