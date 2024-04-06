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

func (u *CartUsecase) GetAllCartItems(ctx context.Context, uID uint) ([]models.CartItem, error) {
	items, err := u.cartRepo.GetAllCartItems(ctx, uID)
	if errors.Is(err, models.ErrEmptyCart) {
		return nil, err
	}
	return items, nil
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
