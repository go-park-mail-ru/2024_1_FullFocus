package usecase

import (
	"context"
	"log"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type OrderUsecase struct {
	orderRepo repository.Orders
	cartRepo  repository.Carts
}

func NewOrderUsecase(or repository.Orders, cr repository.Carts) *OrderUsecase {
	return &OrderUsecase{
		orderRepo: or,
		cartRepo:  cr,
	}
}

func (u *OrderUsecase) Create(ctx context.Context, input models.CreateOrderInput) (uint, error) {
	var orderItems []models.OrderItem
	if input.FromCart {
		log.Printf("from cart")
		cartItems, err := u.cartRepo.GetAllCartItemsID(ctx, input.UserID)
		if err != nil {
			return 0, err
		}
		orderItems = models.ConvertCartItemsToOrderItems(cartItems)
	} else {
		orderItems = input.Items
	}
	orderID, err := u.orderRepo.Create(ctx, input.UserID, orderItems)
	if err != nil {
		return 0, err
	}
	if input.FromCart {
		return orderID, u.cartRepo.DeleteAllCartItems(ctx, input.UserID)
	}
	return orderID, nil
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
