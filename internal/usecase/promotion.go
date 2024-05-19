package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"slices"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/promotion"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type PromotionCache interface {
	Get(ctx context.Context, productID uint) (product models.CachePromoProduct, found bool)
	Set(ctx context.Context, productID uint, product models.CachePromoProduct)
}

const (
	defaultPromoProductsAmount = 3
	percentSale                = "percentSale"
	priceSale                  = "priceSale"
	finalPrice                 = "finalPrice"
)

type PromotionUsecase struct {
	productRepo     repository.Products
	promotionClient promotion.PromotionClient
	cache           PromotionCache
	promoProductIDs []uint
}

func NewPromotionUsecase(pr repository.Products, pc promotion.PromotionClient, c PromotionCache) *PromotionUsecase {
	return &PromotionUsecase{
		productRepo:     pr,
		promotionClient: pc,
		cache:           c,
		promoProductIDs: make([]uint, 0),
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
	if len(u.promoProductIDs) == 0 {
		u.promoProductIDs = append(u.promoProductIDs, input.ProductID)
	}
	var i int
	for i = 0; i < len(u.promoProductIDs); i++ {
		if u.promoProductIDs[i] == 0 {
			u.promoProductIDs[i] = input.ProductID
		}
	}
	if i == 0 || i > len(u.promoProductIDs) {
		u.promoProductIDs = append(u.promoProductIDs, input.ProductID)
	}
	return nil
}

func (u *PromotionUsecase) GetPromoProducts(ctx context.Context, amount uint) ([]models.PromoProduct, error) {
	if amount == 0 {
		amount = defaultPromoProductsAmount
	}
	if len(u.promoProductIDs) == 0 {
		return nil, models.ErrProductNotFound
	}
	randomProductIDs := make([]uint, 0, amount)
	for i := 0; i < int(amount) && i < len(u.promoProductIDs); i++ {
		randomIdx := rand.Int() % len(u.promoProductIDs)
		randomID := u.promoProductIDs[randomIdx]
		found := slices.Contains(randomProductIDs, randomID)
		for found {
			randomIdx = rand.Int() % len(u.promoProductIDs)
			randomID = u.promoProductIDs[randomIdx]
		}
		randomProductIDs = append(randomProductIDs, randomID)
	}
	res := make([]models.PromoProduct, 0, amount)
	prIDs := make([]uint, 0)
	for _, id := range randomProductIDs {
		if product, found := u.cache.Get(ctx, id); found {
			res = append(res, product.Product)
		} else {
			prIDs = append(prIDs, id)
		}
	}
	promoInfo, err := u.promotionClient.GetPromoProductsInfoByIDs(ctx, prIDs)
	if err != nil {
		return nil, err
	}
	productsData, err := u.productRepo.GetProductsByIDs(ctx, prIDs)
	if err != nil {
		return nil, err
	}
	for i, product := range productsData {
		var newPrice uint
		switch promoInfo[i].BenefitType {
		case percentSale:
			newPrice = product.Price / 100 * (100 - promoInfo[i].BenefitValue)
		case priceSale:
			newPrice = product.Price - promoInfo[i].BenefitValue
		case finalPrice:
			newPrice = promoInfo[i].BenefitValue
		}
		promoProduct := models.PromoProduct{
			ProductData:  product,
			BenefitType:  promoInfo[i].BenefitType,
			BenefitValue: promoInfo[i].BenefitValue,
			NewPrice:     newPrice,
		}
		res = append(res, promoProduct)
		u.cache.Set(ctx, product.ID, models.CachePromoProduct{
			Product: promoProduct,
			Empty:   false,
		})
	}
	return res, nil
}

func (u *PromotionUsecase) DeletePromoProduct(ctx context.Context, productID uint) error {
	if err := u.productRepo.MarkProduct(ctx, productID, false); err != nil {
		return err
	}
	if err := u.promotionClient.DeletePromoProductInfo(ctx, productID); err != nil {
		return err
	}
	u.cache.Set(ctx, productID, models.CachePromoProduct{
		Empty: true,
	})
	u.promoProductIDs[slices.Index(u.promoProductIDs, productID)] = 0
	return nil
}
