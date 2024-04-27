package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	grpc "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/delivery/gen"
)

type CsatUsecase struct {
	client grpc.CSATClient
}

func NewCsatUsecase(c grpc.CSATClient) *CsatUsecase {
	return &CsatUsecase{
		client: c,
	}
}

func (u *CsatUsecase) CreatePollRate(ctx context.Context, input models.CreatePollRateInput) error {
	_, err := u.client.CreatePollRate(ctx, &grpc.CreatePollRateRequest{
		ProfileID: uint32(input.ProfileID),
		PollID:    uint32(input.PollID),
		Rate:      uint32(input.Rate),
	})
	return err
}

func (u *CsatUsecase) GetPolls(ctx context.Context, input models.GetPollsInput) ([]models.Poll, error) {
	res, err := u.client.GetPolls(ctx, &grpc.GetPollsRequest{
		ProfileID: uint32(input.ProfileID),
	})
	if err != nil {
		return nil, err
	}
	var polls []models.Poll
	for _, p := range res.Polls {
		polls = append(polls, models.Poll{
			ID:    uint(p.Id),
			Title: p.Title,
			Voted: p.Voted,
		})
	}
	return polls, nil
}

func (u *CsatUsecase) GetPollStats(ctx context.Context, pollID uint) (models.PollStats, error) {
	res, err := u.client.GetPollStats(ctx, &grpc.GetPollStatsRequest{
		PollID: uint32(pollID),
	})
	if err != nil {
		return models.PollStats{}, err
	}
	basicStats := models.StatsData{
		Rates:   res.Stats.Rates,
		Amount:  res.Stats.Amount,
		Above70: res.Stats.Above70,
	}
	primaryStats := models.StatsData{
		Rates:   res.PrimaryStats.Rates,
		Amount:  res.PrimaryStats.Amount,
		Above70: res.PrimaryStats.Above70,
	}
	stats := models.PollStats{
		Stats:        basicStats,
		PrimaryStats: primaryStats,
	}
	return stats, nil
}
