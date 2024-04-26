package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type SortTypePayload struct {
	ID   uint   `json:"sortID"`
	Name string `json:"sortName"`
}

func ConvertSortTypesToDTO(mm []models.SortType) []SortTypePayload {
	sortTypes := make([]SortTypePayload, 0)
	for _, m := range mm {
		sortTypes = append(sortTypes, SortTypePayload{
			ID:   m.ID,
			Name: m.Name,
		})
	}
	return sortTypes
}
