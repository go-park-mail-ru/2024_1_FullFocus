package models

type OrderProduct struct {
	ID          uint
	ProductName string
	Price       uint
	Count       uint
	ImgSrc      string
}

type OrderItem struct {
	ProductID   uint
	Count       uint
	ActualPrice uint
}

func ConvertCartItemsToOrderItems(cartItems []CartItem) []OrderItem {
	orderItems := make([]OrderItem, 0, len(cartItems))
	for _, item := range cartItems {
		orderItems = append(orderItems, OrderItem{
			ProductID:   item.ProductID,
			Count:       item.Count,
			ActualPrice: item.Price,
		})
	}
	return orderItems
}

type Order struct {
	ID         uint
	Sum        uint
	Status     string
	ItemsCount uint
	CreatedAt  string
}

// Create

type CreateOrderInput struct {
	UserID      uint
	Items       []OrderItem
	PromocodeID uint
	FromCart    bool
}

type CreateOrderPayload struct {
	OrderID        uint
	NewPromocodeID uint
}

// Get

type GetOrderPayload struct {
	Products   []OrderProduct
	Sum        uint
	Status     string
	ItemsCount uint
	CreatedAt  string
}

// Update

type UpdateOrderStatusInput struct {
	OrderID   uint
	NewStatus string
}

type UpdateOrderStatusPayload struct {
	OrderID   uint
	OldStatus string
	NewStatus string
}
