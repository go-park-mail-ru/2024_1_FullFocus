package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/promotion"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

const (
	_defaultProductSortType = 0
)

type ProductUsecase struct {
	productRepo     repository.Products
	categoryRepo    repository.Categories
	promotionClient promotion.PromotionClient
}

func NewProductUsecase(pr repository.Products, cr repository.Categories, pc promotion.PromotionClient) *ProductUsecase {
	return &ProductUsecase{
		productRepo:     pr,
		categoryRepo:    cr,
		promotionClient: pc,
	}
}

func (u *ProductUsecase) GetAllProductCards(ctx context.Context, input models.GetAllProductsInput) ([]models.ProductCard, error) {
	if input.PageNum <= 0 || input.PageSize <= 0 {
		return []models.ProductCard{}, models.ErrInvalidParameters
	}
	sorting, err := validateProductSorting(input.Sorting)
	if err != nil {
		return []models.ProductCard{}, err
	}
	input.Sorting = sorting
	return u.productRepo.GetAllProductCards(ctx, input)
}

func (u *ProductUsecase) GetProductByID(ctx context.Context, profileID uint, productID uint) (models.PromoProduct, error) {
	productData, err := u.productRepo.GetProductByID(ctx, profileID, productID)
	if err != nil {
		return models.PromoProduct{}, err
	}
	if productData.OnSale {
		discountData, err := u.promotionClient.GetPromoProductInfoByID(ctx, productID)
		if err != nil {
			return models.PromoProduct{}, err
		}
		newPrice := CalculateDiscountPrice(discountData.BenefitType, discountData.BenefitValue, productData.Price)
		return models.PromoProduct{
			ProductData:  productData,
			BenefitType:  discountData.BenefitType,
			BenefitValue: discountData.BenefitValue,
			NewPrice:     newPrice,
		}, nil
	}
	return models.PromoProduct{
		ProductData: productData,
	}, nil
}

func (u *ProductUsecase) GetProductsByCategoryID(ctx context.Context, input models.GetProductsByCategoryIDInput) (models.GetProductsByCategoryIDPayload, error) {
	sorting, err := validateProductSorting(input.Sorting)
	if err != nil {
		return models.GetProductsByCategoryIDPayload{}, err
	}
	input.Sorting = sorting
	products, err := u.productRepo.GetProductsByCategoryID(ctx, input)
	if err != nil {
		return models.GetProductsByCategoryIDPayload{}, err
	}
	categoryName, err := u.categoryRepo.GetCategoryNameById(ctx, input.CategoryID)
	if err != nil {
		return models.GetProductsByCategoryIDPayload{}, err
	}
	return models.GetProductsByCategoryIDPayload{
		CategoryName: categoryName,
		Products:     products,
	}, nil
}

func (u *ProductUsecase) GetProductsByQuery(ctx context.Context, input models.GetProductsByQueryInput) ([]models.ProductCard, error) {
	sorting, err := validateProductSorting(input.Sorting)
	if err != nil {
		return nil, err
	}
	input.Sorting = sorting
	return u.productRepo.GetProductsByQuery(ctx, input)
}

func validateProductSorting(input models.SortType) (models.SortType, error) {
	if input.ID > 3 {
		defaultSorting, err := helper.GetSortTypeByID(_defaultProductSortType)
		if err != nil {
			return models.SortType{}, models.ErrInternal
		}
		return defaultSorting, nil
	}
	return input, nil
}
