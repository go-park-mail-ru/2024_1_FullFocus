package models

type SortType struct {
	ID        uint
	Name      string
	QueryPart string
}

var Sorting = [...]SortType{
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

const (
	DefaultProductSortType = 0
	DefaultReviewSortType  = 4
)
