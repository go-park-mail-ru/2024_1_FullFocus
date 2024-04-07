package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type OrderUsecase struct {
	cartRepo  repository.Carts
	orderRepo repository.Orders
}

func NewOrderUsecase(cr repository.Carts, or repository.Orders) *OrderUsecase {
	return &OrderUsecase{
		cartRepo:  cr,
		orderRepo: or,
	}
}

func (u *OrderUsecase) Create(ctx context.Context, input models.CreateOrderInput) (uint, error) {
	var orderItems []models.OrderItem
	if input.FromCart {
		// TODO: go to cart and get all Items { productID, amount }
	} else {
		orderItems = input.Items
	}
	return u.orderRepo.Create(ctx, input.UserID, orderItems)
}

func (u *OrderUsecase) GetOrderProducts(ctx context.Context, profileID uint, orderingID uint) ([]models.OrderProduct, error) {
	trueProfileID, err := u.orderRepo.GetProfileIDByOrderingID(ctx, orderingID)
	if err != nil {
		return nil, err
	}
	if profileID != trueProfileID {
		return nil, models.ErrNoAccess
	}
	return u.orderRepo.GetOrderProducts(ctx, orderingID)
}

func (u *OrderUsecase) Delete(ctx context.Context, profileID uint, orderingID uint) error {
	trueProfileID, err := u.orderRepo.GetProfileIDByOrderingID(ctx, orderingID)
	if err != nil {
		return err
	}
	if profileID != trueProfileID {
		return models.ErrNoAccess
	}
	return u.orderRepo.Delete(ctx, orderingID)
}
