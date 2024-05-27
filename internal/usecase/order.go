package usecase

import (
	"context"
	"slices"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/promotion"
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
	promotionClient  promotion.PromotionClient
}

func NewOrderUsecase(or repository.Orders, cr repository.Carts, pr repository.Products, pcr repository.Promocodes, nr repository.Notifications, pc promotion.PromotionClient) *OrderUsecase {
	return &OrderUsecase{
		orderRepo:        or,
		cartRepo:         cr,
		productRepo:      pr,
		promocodeRepo:    pcr,
		notificationRepo: nr,
		promotionClient:  pc,
	}
}

func (u *OrderUsecase) Create(ctx context.Context, input models.CreateOrderInput) (models.CreateOrderPayload, error) {
	var orderItems []models.OrderItem
	if input.FromCart {
		cartItems, err := u.cartRepo.GetAllCartItemsInfo(ctx, input.UserID)
		if err != nil {
			return models.CreateOrderPayload{}, err
		}
		promoProductsIDs := make([]uint, 0)
		for _, product := range cartItems {
			if product.OnSale {
				promoProductsIDs = append(promoProductsIDs, product.ProductID)
			}
		}
		if len(promoProductsIDs) != 0 {
			promoData, err := u.promotionClient.GetPromoProductsInfoByIDs(ctx, promoProductsIDs)
			if err != nil {
				return models.CreateOrderPayload{}, nil
			}
			for i, product := range cartItems {
				if product.OnSale {
					idx := slices.Index(promoProductsIDs, product.ProductID)
					if idx == -1 {
						return models.CreateOrderPayload{}, models.ErrInternal
					}
					cartItems[i].Price = CalculateDiscountPrice(promoData[idx].BenefitType, promoData[idx].BenefitValue, product.Price)
				}
			}
		}
		orderItems = models.ConvertCartItemsToOrderItems(cartItems)
	} else {
		productIDs := make([]uint, 0, len(input.Items))
		for _, item := range input.Items {
			productIDs = append(productIDs, item.ProductID)
		}
		productsData, err := u.productRepo.GetProductsByIDs(ctx, input.UserID, productIDs)
		if err != nil {
			return models.CreateOrderPayload{}, err
		}
		promoProductsIDs := make([]uint, 0)
		for _, product := range productsData {
			if product.OnSale {
				promoProductsIDs = append(promoProductsIDs, product.ID)
			}
		}
		if len(promoProductsIDs) != 0 {
			promoData, err := u.promotionClient.GetPromoProductsInfoByIDs(ctx, promoProductsIDs)
			if err != nil {
				return models.CreateOrderPayload{}, nil
			}
			for i, product := range productsData {
				if product.OnSale {
					idx := slices.Index(promoProductsIDs, product.ID)
					if idx == -1 {
						return models.CreateOrderPayload{}, models.ErrInternal
					}
					productsData[i].Price = CalculateDiscountPrice(promoData[idx].BenefitType, promoData[idx].BenefitValue, product.Price)
				}
			}
		}
		for i, product := range productsData {
			orderItems = append(orderItems, models.OrderItem{
				ProductID:   product.ID,
				Count:       input.Items[i].Count,
				ActualPrice: product.Price,
			})
		}
	}
	var (
		sum uint
		err error
	)
	for _, item := range orderItems {
		sum += item.ActualPrice * item.Count
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

func (u *OrderUsecase) UpdateStatus(ctx context.Context, input models.UpdateOrderStatusInput) (models.UpdateOrderStatusPayload, error) {
	prevStatus, err := u.orderRepo.UpdateStatus(ctx, input.OrderID, input.NewStatus)
	if err != nil {
		return models.UpdateOrderStatusPayload{}, err
	}
	return models.UpdateOrderStatusPayload{
		OrderID:   input.OrderID,
		OldStatus: prevStatus,
		NewStatus: input.NewStatus,
	}, nil
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
