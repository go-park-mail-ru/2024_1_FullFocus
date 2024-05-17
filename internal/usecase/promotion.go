package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/promotion"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

const defaultPromoProductsAmount = 3

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

func (u *PromotionUsecase) GetPromoProducts(ctx context.Context, amount uint) ([]models.PromoProduct, error) {
	if amount == 0 {
		amount = defaultPromoProductsAmount
	}
	promoInfo, err := u.promotionClient.GetPromoProducts(ctx, amount)
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
		res[i].Deadline = promoInfo[i].Deadline
	}
	return res, nil
}
