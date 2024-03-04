package usecase

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type ProductsUsecase struct {
	prodRepo repository.Products
}

func NewProductsUsecase(pr repository.Products) *ProductsUsecase {
	return &ProductsUsecase{
		prodRepo: pr,
	}
}

func (u *ProductsUsecase) GetProducts(lastID, limit int) ([]models.Product, error) {
	prods, err := u.prodRepo.GetProducts(lastID, limit)
	if err != nil {
		return nil, err
	}
	return prods, nil
}
