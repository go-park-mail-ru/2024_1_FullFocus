package models

type Product struct {
	PrID        uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Img         string `json:"img-link"`
}
