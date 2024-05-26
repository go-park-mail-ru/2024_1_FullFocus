package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/pkg/errors"
)

const _activationStringLen = 6

type OrderUsecase struct {
	orderRepo        repository.Orders
	cartRepo         repository.Carts
	productRepo      repository.Products
	promocodeRepo    repository.Promocodes
	notificationRepo repository.Notifications
}

func NewOrderUsecase(or repository.Orders, cr repository.Carts, pr repository.Products, pcr repository.Promocodes, nr repository.Notifications) *OrderUsecase {
	return &OrderUsecase{
		orderRepo:        or,
		cartRepo:         cr,
		productRepo:      pr,
		promocodeRepo:    pcr,
		notificationRepo: nr,
	}
}

func (u *OrderUsecase) Create(ctx context.Context, input models.CreateOrderInput) (models.CreateOrderPayload, error) {
	var orderItems []models.OrderItem
	if input.FromCart {
		cartItems, err := u.cartRepo.GetAllCartItemsID(ctx, input.UserID)
		if err != nil {
			return models.CreateOrderPayload{}, err
		}
		orderItems = models.ConvertCartItemsToOrderItems(cartItems)
	} else {
		orderItems = input.Items
	}
	sum, err := u.productRepo.GetTotalPrice(ctx, orderItems)
	if err != nil {
		return models.CreateOrderPayload{}, err
	}
	var promoUsed bool
	if input.PromocodeID != 0 {
		info := models.ApplyPromocodeInput{
			Sum:       sum,
			PromoID:   input.PromocodeID,
			ProfileID: input.UserID,
		}
		sum, err = u.promocodeRepo.ApplyPromocode(ctx, info)
		if err != nil {
			return models.CreateOrderPayload{}, err
		}
		promoUsed = true
	}
	orderID, err := u.orderRepo.Create(ctx, input.UserID, sum, orderItems)
	if err != nil {
		return models.CreateOrderPayload{}, err
	}
	result := models.CreateOrderPayload{
		OrderID: orderID,
	}
	if promoUsed {
		if err = u.promocodeRepo.DeleteUsedPromocode(ctx, input.PromocodeID); err != nil {
			return models.CreateOrderPayload{}, err
		}
	}
	availablePromoID, err := u.promocodeRepo.GetNewPromocode(ctx, sum)
	if err == nil {
		promoInfo := models.CreatePromocodeItemInput{
			PromocodeID: availablePromoID,
			ProfileID:   input.UserID,
			Code:        helper.RandActivationCode(_activationStringLen),
		}
		if err = u.promocodeRepo.CreatePromocodeItem(ctx, promoInfo); err != nil {
			return result, err
		}
		result.NewPromocodeID = availablePromoID
	} else if !errors.Is(err, models.ErrNoPromocode) {
		return models.CreateOrderPayload{}, err
	}
	if input.FromCart {
		return result, u.cartRepo.DeleteAllCartItems(ctx, input.UserID)
	}
	return result, nil
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

func (u *OrderUsecase) UpdateStatus(ctx context.Context, input models.UpdateOrderStatusInput) error {
	profileID, err := u.orderRepo.GetProfileIDByOrderID(ctx, input.OrderID)
	if err != nil {
		return err
	}
	prevStatus, err := u.orderRepo.UpdateStatus(ctx, input.OrderID, input.NewStatus)
	if err != nil {
		return err
	}
	payload := fmt.Sprintf(`{
		"type": "orderStatusChange",
		"data": {
			  "orderID": %d,
			  "oldStatus": "%s",
			  "newStatus": "%s"
		 }
	}`, input.OrderID, prevStatus, input.NewStatus)
	notification := models.CreateNotificationInput{
		Type:    "order_status_change",
		Payload: payload,
	}
	if err = u.notificationRepo.CreateNotification(ctx, profileID, notification); err != nil {
		return err
	}
	return u.notificationRepo.SendNotification(ctx, profileID, payload) // for now does nothing
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
