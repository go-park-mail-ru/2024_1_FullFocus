package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type OrderUsecase struct {
	// cartRepo  repository.Carts
	orderRepo repository.Orders
}

func NewOrderUsecase( /*cr repository.Carts, */ or repository.Orders) *OrderUsecase {
	return &OrderUsecase{
		// cartRepo:  cr,
		orderRepo: or,
	}
}

func (u *OrderUsecase) Create(ctx context.Context, input models.CreateOrderInput) (uint, error) {
	var orderItems []models.OrderItem
	// if input.FromCart {
	// TODO: go to cart and get all Items { productID, amount }
	// } else {
	if !input.FromCart {
		orderItems = input.Items
	}
	// }
	return u.orderRepo.Create(ctx, input.UserID, orderItems)
}

func (u *OrderUsecase) GetOrderByID(ctx context.Context, profileID uint, orderID uint) (models.GetOrderPayload, error) {
	trueProfileID, err := u.orderRepo.GetProfileIDByOrderID(ctx, orderID)
	if err != nil {
		return models.GetOrderPayload{}, err
	}
	if profileID != trueProfileID {
		return models.GetOrderPayload{}, models.ErrNoAccess
	}
	return u.orderRepo.GetOrderByID(ctx, orderID)
}

func (u *OrderUsecase) GetAllOrders(ctx context.Context, profileID uint) ([]models.Order, error) {
	return u.orderRepo.GetAllOrders(ctx, profileID)
}

func (u *OrderUsecase) Delete(ctx context.Context, profileID uint, orderID uint) error {
	trueProfileID, err := u.orderRepo.GetProfileIDByOrderID(ctx, orderID)
	if err != nil {
		return err
	}
	if profileID != trueProfileID {
		return models.ErrNoAccess
	}
	return u.orderRepo.Delete(ctx, orderID)
}
