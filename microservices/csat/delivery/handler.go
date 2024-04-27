package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/delivery/gen"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/models"
	"github.com/golang/protobuf/ptypes/empty"
)

type CSATs interface {
	GetAllPolls(ctx context.Context) ([]models.Poll, error)
	CreatePollRate(ctx context.Context, input models.CreatePollRate) error
}

type CSATHandler struct {
	CSATUsecase CSATs
	CSATServer  gen.CSATServer
}

func NewCSATHandler(u CSATs) *CSATHandler {
	return &CSATHandler{
		CSATUsecase: u,
	}
}

func (h *CSATHandler) GetPolls(ctx context.Context, r *empty.Empty) (*gen.GetPollsResponse, error) {
	polls, err := h.CSATUsecase.GetAllPolls(ctx)
	if err != nil {
		return &gen.GetPollsResponse{}, err
	}
	return &gen.GetPollsResponse{
		Polls: polls,
	}, nil
}
