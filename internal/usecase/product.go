package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
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
	return u.productRepo.GetAllProductCards(ctx, input)
}

func (u *ProductUsecase) GetProductById(ctx context.Context, profileID uint, productID uint) (models.Product, error) {
	return u.productRepo.GetProductById(ctx, profileID, productID)
}

func (u *ProductUsecase) GetProductsByCategoryId(ctx context.Context, input models.GetProductsByCategoryIDInput) ([]models.ProductCard, error) {
	return u.productRepo.GetProductsByCategoryId(ctx, input)
}
