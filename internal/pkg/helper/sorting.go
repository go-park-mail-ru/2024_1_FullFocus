package helper

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

var sorting = [...]models.SortType{
	{
		ID:   0,
		Name: "Не сортировать",
	},
	{
		ID:        1,
		Name:      "Сначала дорогие",
		QueryPart: "price DESC",
	},
	{
		ID:        2,
		Name:      "Сначала недорогие",
		QueryPart: "price ASC",
	},
	{
		ID:        3,
		Name:      "Сначала с лучшей оценкой",
		QueryPart: "rating DESC",
	},
	{
		ID:        4,
		Name:      "Сначала новые",
		QueryPart: "created_at DESC",
	},
}

func GetSortParams(r *http.Request) (models.SortType, error) {
	sortID, _ := strconv.Atoi(r.URL.Query().Get("sortID"))
	if sortID < 0 || sortID > len(sorting)-1 {
		return models.SortType{}, models.ErrInvalidField
	}
	return sorting[sortID], nil
}

func GetSortTypeByID(ID int) (models.SortType, error) {
	if ID < 0 || ID > len(sorting)-1 {
		return models.SortType{}, models.ErrInvalidParameters
	}
	return sorting[ID], nil
}

func GetProductSortTypes() []models.SortType {
	return sorting[1:4]
}

func GetReviewSortTypes() []models.SortType {
	return sorting[3:]
}

func ApplySorting(q string, sorting string) string {
	if sorting != "" {
		return fmt.Sprintf(q, "ORDER BY "+sorting)
	}
	return fmt.Sprintf(q, "")
}
