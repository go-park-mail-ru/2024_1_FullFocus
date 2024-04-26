package dao

type ProductSuggest struct {
	Name string `json:"product_name"`
}

type CategorySuggest struct {
	ID   uint   `json:"category_id"`
	Name string `json:"category_name"`
}
