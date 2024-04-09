package usecase

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type CartUsecase struct {
	cartRepo repository.Carts
}

func NewCartUsecase(cr repository.Carts) *CartUsecase {
	return &CartUsecase{
		cartRepo: cr,
	}
}

func (u *CartUsecase) GetAllCartItems(ctx context.Context, uID uint) (models.CartContent, error) {
	products, err := u.cartRepo.GetAllCartItems(ctx, uID)
	if errors.Is(err, models.ErrEmptyCart) {
		return models.CartContent{}, err
	}

	var sum, count uint
	for i, product := range products {
		products[i].Cost = product.Price * product.Count
		count += product.Count
		sum += product.Cost
	}

	content := models.CartContent{
		Products:   products,
		TotalCount: count,
		TotalCost:  sum,
	}
	return content, nil
}

func (u *CartUsecase) UpdateCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	newCount, err := u.cartRepo.UpdateCartItem(ctx, uID, prID)
	if errors.Is(err, models.ErrNoProduct) {
		return 0, err
	}
	return newCount, nil
}

func (u *CartUsecase) DeleteCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	newCount, err := u.cartRepo.DeleteCartItem(ctx, uID, prID)
	if errors.Is(err, models.ErrNoProduct) {
		return 0, err
	}
	return newCount, nil
}

func (u *CartUsecase) DeleteAllCartItems(ctx context.Context, uID uint) error {
	return u.cartRepo.DeleteAllCartItems(ctx, uID)
}
