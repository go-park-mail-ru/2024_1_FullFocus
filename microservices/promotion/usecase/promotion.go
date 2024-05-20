package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/models"
)

type Promotion interface {
	CreatePromoProductInfo(ctx context.Context, input models.PromoData) error
	GetAllPromoProductsIDs(ctx context.Context) ([]uint, error)
	GetPromoProductsInfoByIDs(ctx context.Context, prIDs []uint) ([]models.PromoData, error)
	DeletePromoProductInfo(ctx context.Context, prID uint) error
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

func (u *Usecase) GetAllPromoProductsIDs(ctx context.Context) ([]uint, error) {
	return u.repo.GetAllPromoProductsIDs(ctx)
}

func (u *Usecase) GetPromoProductsInfoByIDs(ctx context.Context, prIDs []uint) ([]models.PromoData, error) {
	return u.repo.GetPromoProductsInfoByIDs(ctx, prIDs)
}

func (u *Usecase) DeletePromoProductInfo(ctx context.Context, prID uint) error {
	return u.repo.DeletePromoProductInfo(ctx, prID)
}
