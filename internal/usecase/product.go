package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

const (
	_defaultProductSortType = 0
)

type ProductUsecase struct {
	productRepo repository.Products
}

func NewProductUsecase(pr repository.Products) *ProductUsecase {
	return &ProductUsecase{
		productRepo: pr,
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

func (u *ProductUsecase) GetProductByID(ctx context.Context, profileID uint, productID uint) (models.Product, error) {
	return u.productRepo.GetProductByID(ctx, profileID, productID)
}

func (u *ProductUsecase) GetProductsByCategoryID(ctx context.Context, input models.GetProductsByCategoryIDInput) ([]models.ProductCard, error) {
	sorting, err := validateProductSorting(input.Sorting)
	if err != nil {
		return []models.ProductCard{}, err
	}
	input.Sorting = sorting
	return u.productRepo.GetProductsByCategoryID(ctx, input)
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
