package helper

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

func GetSortParams(r *http.Request, product bool) models.SortType {
	var sortID int
	if sortIDstr := r.URL.Query().Get("sortID"); sortIDstr != "" {
		if product {
			sortID = models.DefaultProductSortType
		} else {
			sortID = models.DefaultReviewSortType
		}
	}
	return models.Sorting[sortID]
}

func GetProductSortTypes() []models.SortType {
	return models.Sorting[1:4]
}

func GetReviewSortTypes() []models.SortType {
	return models.Sorting[3:]
}
