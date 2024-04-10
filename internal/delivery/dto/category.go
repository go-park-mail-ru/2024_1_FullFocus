package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ConvertCategoriesToDto(mm []models.Category) []Category {
	categories := make([]Category, 0)
	for _, m := range mm {
		categories = append(categories, Category{
			ID:   m.ID,
			Name: m.Name,
		})
	}
	return categories
}
