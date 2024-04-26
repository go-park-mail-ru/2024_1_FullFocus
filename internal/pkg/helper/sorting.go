package helper

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

func GetSortParams(r *http.Request, product bool) models.SortType {
	sortID, err := strconv.Atoi(r.URL.Query().Get("sortID"))
	if err != nil || sortID < 0 || sortID > 4 {
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

func ApplySorting(q string, sorting models.SortType) string {
	if sorting.ID != 0 {
		return fmt.Sprintf(q, "ORDER BY "+sorting.QueryPart)
	}
	return fmt.Sprintf(q, "")
}
