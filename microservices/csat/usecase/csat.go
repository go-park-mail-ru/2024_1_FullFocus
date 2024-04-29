package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/models"
)

type csats interface {
	GetAllPolls(context.Context, uint) ([]models.Poll, error)
	CreatePollRate(context.Context, models.CreatePollRate) error
	GetPollStats(context.Context, uint) (string, []models.StatRate, error)
}

type CSATUsecase struct {
	csatRepo csats
}

func NewCSATUsecase(r csats) *CSATUsecase {
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

func (u *CSATUsecase) GetPollStats(ctx context.Context, pollID uint) (models.PollStats, error) {
	title, statRates, err := u.csatRepo.GetPollStats(ctx, pollID)
	if err != nil {
		return models.PollStats{}, err
	}

	rates := make([]uint, 10)
	var amount, above70 uint
	for _, rate := range statRates {
		amount += rate.Amount
		if rate.Rate >= 7 {
			above70 += rate.Amount
		}
		rates[rate.Rate-1] += rate.Amount
	}

	return models.PollStats{
		PollTitle: title,
		Stats: models.StatData{
			Amount:  amount,
			Above70: above70,
			Rates:   rates,
		},
	}, nil
}
