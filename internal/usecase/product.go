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

func (u *ProductUsecase) GetAllProductCards(ctx context.Context, pageNum uint, perPage uint) ([]models.ProductCard, error) {
	if pageNum <= 0 || perPage <= 0 {
		return []models.ProductCard{}, models.ErrInvalidParameters
	}
	return u.productRepo.GetAllProductCards(ctx, pageNum, perPage)
}

func (u *ProductUsecase) GetProductById(ctx context.Context, productID uint) (models.Product, error) {
	return u.productRepo.GetProductById(ctx, productID)
}
