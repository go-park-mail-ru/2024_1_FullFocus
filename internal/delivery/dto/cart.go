package dto

type CartProduct struct {
	ProductID uint   `json:"productId"`
	Name      string `json:"name"`
	Price     uint   `json:"price"`
	Count     uint   `json:"count"`
	Img       string `json:"imgsrc"`
}

type CartItem struct {
	ProductId uint `json:"productId"`
	Count     uint `json:"count"`
}
