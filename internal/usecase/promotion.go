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
	Get(ctx context.Context, productID uint) (product models.PromoProduct, found bool)
	Set(ctx context.Context, productID uint, product models.PromoProduct)
	Remove(ctx context.Context, productID uint)
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

func NewPromotionUsecase(ctx context.Context, pr repository.Products, pc promotion.PromotionClient, c PromotionCache) *PromotionUsecase {
	u := &PromotionUsecase{
		productRepo:     pr,
		promotionClient: pc,
		cache:           c,
		promoProductIDs: make([]uint, 0),
	}
	avaliablePrIDs, _ := u.promotionClient.GetAllPromoProductsIDs(ctx)
	if len(avaliablePrIDs) != 0 {
		u.promoProductIDs = avaliablePrIDs
	}
	return u
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
	var i int
	for i = 0; i < len(u.promoProductIDs); i++ {
		if u.promoProductIDs[i] == 0 {
			u.promoProductIDs[i] = input.ProductID
			return nil
		}
	}
	if i == 0 || i >= len(u.promoProductIDs) {
		u.promoProductIDs = append(u.promoProductIDs, input.ProductID)
	}
	return nil
}

func (u *PromotionUsecase) GetPromoProductInfoByID(ctx context.Context, productID uint, profileID uint) (models.PromoProduct, error) {
	if !slices.Contains(u.promoProductIDs, productID) {
		return models.PromoProduct{}, models.ErrNoProduct
	}
	if promoData, found := u.cache.Get(ctx, productID); found {
		return promoData, nil
	}
	promoData, err := u.promotionClient.GetPromoProductInfoByID(ctx, productID)
	if err != nil {
		return models.PromoProduct{}, err
	}
	productData, err := u.productRepo.GetProductByID(ctx, profileID, productID)
	if err != nil {
		return models.PromoProduct{}, err
	}
	newPrice := CalculateDiscountPrice(promoData.BenefitType, promoData.BenefitValue, productData.Price)
	promoProduct := models.PromoProduct{
		ProductData:  productData,
		BenefitType:  promoData.BenefitType,
		BenefitValue: promoData.BenefitValue,
		NewPrice:     newPrice,
	}
	u.cache.Set(ctx, productID, promoProduct)
	return promoProduct, nil
}

func (u *PromotionUsecase) GetPromoProductCards(ctx context.Context, amount uint, profileID uint) ([]models.PromoProductCard, error) {
	if amount == 0 {
		amount = defaultPromoProductsAmount
	}
	if u.getAvailiablePromoProductsCount() == 0 {
		return nil, models.ErrNoProduct
	}
	randomProductIDs := make([]uint, 0, amount)
	for i := 0; i < int(amount) && i < u.getAvailiablePromoProductsCount(); i++ {
		var (
			found    bool = true
			randomID uint
		)
		for found {
			randomIdx := rand.Int() % len(u.promoProductIDs)
			randomID = u.promoProductIDs[randomIdx]
			if randomID == 0 {
				continue
			}
			found = slices.Contains(randomProductIDs, randomID)
		}
		randomProductIDs = append(randomProductIDs, randomID)
	}
	res := make([]models.PromoProductCard, 0, amount)
	prIDs := make([]uint, 0)
	for _, id := range randomProductIDs {
		if product, found := u.cache.Get(ctx, id); found {
			res = append(res, models.ConvertPromoProductToCard(product))
		} else {
			prIDs = append(prIDs, id)
		}
	}
	if len(prIDs) != 0 {
		promoInfo, err := u.promotionClient.GetPromoProductsInfoByIDs(ctx, prIDs)
		if err != nil {
			return nil, err
		}
		productsData, err := u.productRepo.GetProductsByIDs(ctx, profileID, prIDs)
		if err != nil {
			return nil, err
		}
		for i, product := range productsData {
			newPrice := CalculateDiscountPrice(promoInfo[i].BenefitType, promoInfo[i].BenefitValue, product.Price)
			promoProduct := models.PromoProduct{
				ProductData:  product,
				BenefitType:  promoInfo[i].BenefitType,
				BenefitValue: promoInfo[i].BenefitValue,
				NewPrice:     newPrice,
			}
			res = append(res, models.ConvertPromoProductToCard(promoProduct))
			u.cache.Set(ctx, product.ID, promoProduct)
		}
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
	u.cache.Remove(ctx, productID)
	idx := slices.Index(u.promoProductIDs, productID)
	if idx != -1 {
		u.promoProductIDs[idx] = 0
	}
	return nil
}

func (u *PromotionUsecase) getAvailiablePromoProductsCount() int {
	var count int
	for _, id := range u.promoProductIDs {
		if id != 0 {
			count++
		}
	}
	return count
}

func CalculateDiscountPrice(benefitType string, benefitValue, oldPrice uint) uint {
	var newPrice uint
	switch benefitType {
	case percentSale:
		newPrice = oldPrice / 100 * (100 - benefitValue)
	case priceSale:
		newPrice = oldPrice - benefitValue
	case finalPrice:
		newPrice = benefitValue
	}
	return newPrice
}
