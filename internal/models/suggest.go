package models

type CategorySuggest struct {
	ID   uint
	Name string
}

type Suggestion struct {
	Categories []CategorySuggest
	Products   []string
}
