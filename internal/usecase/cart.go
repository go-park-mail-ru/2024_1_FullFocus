package usecase

import (
	"context"
	"errors"
	"slices"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/promotion"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type CartUsecase struct {
	cartRepo        repository.Carts
	promotionClient promotion.PromotionClient
}

func NewCartUsecase(cr repository.Carts, pc promotion.PromotionClient) *CartUsecase {
	return &CartUsecase{
		cartRepo:        cr,
		promotionClient: pc,
	}
}

func (u *CartUsecase) GetAllCartItems(ctx context.Context, uID uint) (models.CartContent, error) {
	products, err := u.cartRepo.GetAllCartItems(ctx, uID)
	if errors.Is(err, models.ErrEmptyCart) {
		return models.CartContent{}, err
	}
	promoProductsIDs := make([]uint, 0)
	for _, product := range products {
		if product.OnSale {
			promoProductsIDs = append(promoProductsIDs, product.ProductID)
		}
	}
	if len(promoProductsIDs) != 0 {
		promoData, err := u.promotionClient.GetPromoProductsInfoByIDs(ctx, promoProductsIDs)
		if err != nil {
			return models.CartContent{}, nil
		}
		for i, product := range products {
			if product.OnSale {
				idx := slices.Index(promoProductsIDs, product.ProductID)
				if idx == -1 {
					return models.CartContent{}, models.ErrInternal
				}
				products[i].BenefitType = promoData[idx].BenefitType
				products[i].BenefitValue = promoData[idx].BenefitValue
				products[i].NewPrice = CalculateDiscountPrice(promoData[idx].BenefitType, promoData[idx].BenefitValue, product.Price)
			} else {
				products[i].NewPrice = product.Price
			}
		}
	}
	var sum, count uint
	for i, product := range products {
		products[i].Cost = product.NewPrice * product.Count
		count += product.Count
		sum += products[i].Cost
	}
	content := models.CartContent{
		Products:   products,
		TotalCount: count,
		TotalCost:  sum,
	}
	return content, nil
}

func (u *CartUsecase) GetCartItemsAmount(ctx context.Context, uID uint) (uint, error) {
	return u.cartRepo.GetCartItemsAmount(ctx, uID)
}

func (u *CartUsecase) UpdateCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	newCount, err := u.cartRepo.UpdateCartItem(ctx, uID, prID)
	if errors.Is(err, models.ErrNoProduct) {
		return 0, err
	}
	return newCount, nil
}

func (u *CartUsecase) DeleteCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	newCount, err := u.cartRepo.DeleteCartItem(ctx, uID, prID)
	if errors.Is(err, models.ErrNoProduct) {
		return 0, err
	}
	return newCount, nil
}

func (u *CartUsecase) DeleteAllCartItems(ctx context.Context, uID uint) error {
	return u.cartRepo.DeleteAllCartItems(ctx, uID)
}
