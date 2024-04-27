package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/models"
)

type CSATs interface {
	GetAllPolls(context.Context, uint) ([]models.Poll, error)
	CreatePollRate(context.Context, models.CreatePollRate) error
}

type CSATUsecase struct {
	csatRepo CSATs
}

func NewCSATUsecase(r CSATs) *CSATUsecase {
	return &CSATUsecase{
		csatRepo: r,
	}
}

func (u *CSATUsecase) GetAllPolls(ctx context.Context, profileID uint) ([]models.Poll, error) {
	return u.csatRepo.GetAllPolls(ctx, profileID)
}

func (u *CSATUsecase) CreatePollRate(ctx context.Context, input models.CreatePollRate) error {
	return u.csatRepo.CreatePollRate(ctx, input)
}
