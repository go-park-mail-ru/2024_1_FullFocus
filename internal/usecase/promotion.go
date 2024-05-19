package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/promotion"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

const (
	defaultPromoProductsAmount = 3
	percentSale                = "percentSale"
	priceSale                  = "priceSale"
	finalPrice                 = "finalPrice"
)

type PromotionUsecase struct {
	productRepo     repository.Products
	promotionClient promotion.PromotionClient
}

func NewPromotionUsecase(pr repository.Products, pc promotion.PromotionClient) *PromotionUsecase {
	return &PromotionUsecase{
		productRepo:     pr,
		promotionClient: pc,
	}
}

func (u *PromotionUsecase) CreatePromoProduct(ctx context.Context, input models.PromoData) error {
	switch input.BenefitType {
	case percentSale:
		if input.BenefitValue == 0 || input.BenefitValue >= 100 {
			return models.ErrInvalidBenefitValue
		}
	case priceSale:
		oldPrice, err := u.productRepo.GetProductPriceByID(ctx, input.ProductID)
		if err != nil {
			return err
		}
		if input.BenefitValue == 0 || input.BenefitValue >= oldPrice {
			return models.ErrInvalidBenefitValue
		}
	case finalPrice:
		oldPrice, err := u.productRepo.GetProductPriceByID(ctx, input.ProductID)
		if err != nil {
			return err
		}
		if input.BenefitValue == 0 || input.BenefitValue == oldPrice {
			return models.ErrInvalidBenefitValue
		}
	default:
		return helper.NewValidationError("invalid benefit type", fmt.Sprintf("Неправльный тип акции, возможные варианты: %s, %s, %s", percentSale, priceSale, finalPrice))
	}
	if err := u.promotionClient.CreatePromoProductInfo(ctx, input); err != nil {
		return err
	}
	if err := u.productRepo.MarkProduct(ctx, input.ProductID, true); err != nil {
		if err := u.promotionClient.DeletePromoProductInfo(ctx, input.ProductID); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (u *PromotionUsecase) GetPromoProducts(ctx context.Context, amount uint) ([]models.PromoProduct, error) {
	if amount == 0 {
		amount = defaultPromoProductsAmount
	}
	promoInfo, err := u.promotionClient.GetPromoProductsInfo(ctx, amount)
	if err != nil {
		return nil, err
	}
	IDs := make([]uint, 0, len(promoInfo))
	for _, promo := range promoInfo {
		IDs = append(IDs, promo.ProductID)
	}
	productsData, err := u.productRepo.GetProductsByIDs(ctx, IDs)
	if err != nil {
		return nil, err
	}
	res := make([]models.PromoProduct, len(productsData))
	for i, product := range productsData {
		res[i].ProductData = product
		res[i].BenefitType = promoInfo[i].BenefitType
		res[i].BenefitValue = promoInfo[i].BenefitValue
		switch promoInfo[i].BenefitType {
		case percentSale:
			res[i].NewPrice = product.Price / 100 * (100 - promoInfo[i].BenefitValue)
		case priceSale:
			res[i].NewPrice = product.Price - promoInfo[i].BenefitValue
		case finalPrice:
			res[i].NewPrice = promoInfo[i].BenefitValue
		}
	}
	return res, nil
}

func (u *PromotionUsecase) DeletePromoProduct(ctx context.Context, productID uint) error {
	if err := u.productRepo.MarkProduct(ctx, productID, false); err != nil {
		return err
	}
	return u.promotionClient.DeletePromoProductInfo(ctx, productID)
}
