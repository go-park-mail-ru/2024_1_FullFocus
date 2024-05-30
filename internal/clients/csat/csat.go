package csat

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type CsatClient interface {
	GetAllPolls(context.Context, uint) ([]models.Poll, error)
	CreatePollRate(context.Context, models.CreatePollRateInput) error
	GetPollStats(context.Context, uint) (models.PollStats, error)
}
