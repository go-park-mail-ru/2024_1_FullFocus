package usecase

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type ProductUsecase struct {
	prodRepo repository.Products
}

func NewProductUsecase(pr repository.Products) *ProductUsecase {
	return &ProductUsecase{
		prodRepo: pr,
	}
}

func (u *ProductUsecase) GetProducts(lastID, limit int) ([]models.Product, error) {
	prods, err := u.prodRepo.GetProducts(lastID, limit)
	if err != nil {
		return nil, err
	}
	return prods, nil
}
