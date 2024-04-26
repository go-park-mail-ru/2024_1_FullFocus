package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type ProductUsecase struct {
	productRepo  repository.Products
	categoryRepo repository.Categories
}

func NewProductUsecase(pr repository.Products, cr repository.Categories) *ProductUsecase {
	return &ProductUsecase{
		productRepo:  pr,
		categoryRepo: cr,
	}
}

func (u *ProductUsecase) GetAllProductCards(ctx context.Context, input models.GetAllProductsInput) ([]models.ProductCard, error) {
	if input.PageNum <= 0 || input.PageSize <= 0 {
		return []models.ProductCard{}, models.ErrInvalidParameters
	}
	return u.productRepo.GetAllProductCards(ctx, input)
}

func (u *ProductUsecase) GetProductByID(ctx context.Context, profileID uint, productID uint) (models.Product, error) {
	return u.productRepo.GetProductByID(ctx, profileID, productID)
}

func (u *ProductUsecase) GetProductsByCategoryID(ctx context.Context, input models.GetProductsByCategoryIDInput) (models.GetProductsByCategoryIDPayload, error) {
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
	return u.productRepo.GetProductsByQuery(ctx, input)
}
