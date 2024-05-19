package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type PromocodeUsecase struct {
	promocodeRepo repository.Promocodes
}

func NewPromocodeUsecase(pr repository.Promocodes) *PromocodeUsecase {
	return &PromocodeUsecase{
		promocodeRepo: pr,
	}
}

func (u *PromocodeUsecase) GetPromocodeItemByActivationCode(ctx context.Context, pID uint, code string) (models.PromocodeActivationTerms, error) {
	return u.promocodeRepo.GetPromocodeItemByActivationCode(ctx, pID, code)
}

func (u *PromocodeUsecase) GetPromocodeByID(ctx context.Context, promocodeID uint) (models.Promocode, error) {
	return u.promocodeRepo.GetPromocodeByID(ctx, promocodeID)
}

func (u *PromocodeUsecase) GetAvailablePromocodes(ctx context.Context, profileID uint) ([]models.PromocodeItem, error) {
	return u.promocodeRepo.GetAvailablePromocodes(ctx, profileID)
}
