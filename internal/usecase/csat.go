package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/csat"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type CsatUsecase struct {
	client csat.CsatClient
}

func NewCsatUsecase(c csat.CsatClient) *CsatUsecase {
	return &CsatUsecase{
		client: c,
	}
}

func (u *CsatUsecase) CreatePollRate(ctx context.Context, input models.CreatePollRateInput) error {
	return u.client.CreatePollRate(ctx, input)
}

func (u *CsatUsecase) GetPolls(ctx context.Context, userID uint) ([]models.Poll, error) {
	return u.client.GetAllPolls(ctx, userID)
}

func (u *CsatUsecase) GetPollStats(ctx context.Context, pollID uint) (models.PollStats, error) {
	return u.client.GetPollStats(ctx, pollID)
}
