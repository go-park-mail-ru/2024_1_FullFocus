package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type CategorySuggest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Suggestion struct {
	Categories []CategorySuggest `json:"categories"`
	Products   []string          `json:"products"`
}

func convertCategorySuggestToDTO(suggestions []models.CategorySuggest) []CategorySuggest {
	var categorySuggests []CategorySuggest
	for _, suggestion := range suggestions {
		categorySuggests = append(categorySuggests, CategorySuggest{
			ID:   suggestion.ID,
			Name: suggestion.Name,
		})
	}
	return categorySuggests
}

func ConvertSuggestionToDTO(suggestion models.Suggestion) Suggestion {
	return Suggestion{
		Categories: convertCategorySuggestToDTO(suggestion.Categories),
		Products:   suggestion.Products,
	}
}
