package repository

import (
	"sync"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type ProductRepo struct {
	sync.Mutex
	products []models.Product
}

func NewProductRepo() *ProductRepo {
	r := &ProductRepo{
		products: make([]models.Product, 20),
	}
	return r
}

func (r *ProductRepo) GetProducts(lastID, limit int) ([]models.Product, error) {
	r.Lock()
	defer r.Unlock()
	prods := make([]models.Product, limit)
	for i := lastID; i < lastID+limit; i++ {
		prods = append(prods, r.products[i])
	}
	return prods, nil
}
