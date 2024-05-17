package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/models"
)

type Promotion interface {
	CreatePromoProductInfo(ctx context.Context, input models.PromoData) error
	GetPromoProductsInfo(ctx context.Context, amount uint32) ([]models.PromoData, error)
	DeletePromoProductInfo(ctx context.Context, pID uint32) error
}

type Usecase struct {
	repo Promotion
}

func NewUsecase(r Promotion) *Usecase {
	return &Usecase{
		repo: r,
	}
}

func (u *Usecase) CreatePromoProductInfo(ctx context.Context, input models.PromoData) error {
	return u.repo.CreatePromoProductInfo(ctx, input)
}

func (u *Usecase) GetPromoProductsInfo(ctx context.Context, amount uint32) ([]models.PromoData, error) {
	return u.repo.GetPromoProductsInfo(ctx, amount)
}

func (u *Usecase) DeletePromoProductInfo(ctx context.Context, pID uint32) error {
	return u.repo.DeletePromoProductInfo(ctx, pID)
}
