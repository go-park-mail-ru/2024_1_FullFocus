package dao

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type CategoryTable struct {
	ID       uint   `db:"id"`
	Name     string `db:"category_name"`
	ParentID uint   `db:"parent_id"`
}

func ConvertTablesToCategories(tt []CategoryTable) []models.Category {
	categories := make([]models.Category, 0)
	for _, t := range tt {
		categories = append(categories, models.Category{
			ID:   t.ID,
			Name: t.Name,
		})
	}
	return categories
}
