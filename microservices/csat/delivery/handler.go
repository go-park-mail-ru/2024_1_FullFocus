package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/delivery/gen"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/models"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type CSATs interface {
	GetAllPolls(context.Context, uint) ([]models.Poll, error)
	CreatePollRate(context.Context, models.CreatePollRate) error
	GetPollStats(context.Context, uint) (models.PollStats, error)
}

type CSATHandler struct {
	csatUsecase CSATs
	gen.UnimplementedCSATServer
}

func NewCSATHandler(u CSATs) *CSATHandler {
	return &CSATHandler{
		csatUsecase: u,
	}
}

func (h *CSATHandler) GetPolls(ctx context.Context, r *gen.CreatePollRateRequest) (*gen.GetPollsResponse, error) {
	polls, err := h.csatUsecase.GetAllPolls(ctx, uint(r.ProfileID))
	if err != nil {
		return &gen.GetPollsResponse{}, err
	}

	payload := gen.GetPollsResponse{
		Polls: make([]*gen.Poll, 0),
	}
	for _, poll := range polls {
		payload.Polls = append(payload.Polls, &gen.Poll{
			Id:    uint32(poll.ID),
			Title: poll.Title,
			Voted: poll.Voted,
		})
	}
	return &payload, nil
}

func (h *CSATHandler) CreatePollRate(ctx context.Context, r *gen.CreatePollRateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	input := models.CreatePollRate{
		ProfileID: uint(r.ProfileID),
		PollID:    uint(r.PollID),
		Rate:      uint(r.Rate),
	}

	if err := h.csatUsecase.CreatePollRate(ctx, input); err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}

func (h *CSATHandler) GetPollStats(ctx context.Context, r *gen.GetPollStatsRequest) (*gen.GetPollStatsResponse, error) {
	stats, err := h.csatUsecase.GetPollStats(ctx, uint(r.PollID))
	if err != nil {
		return &gen.GetPollStatsResponse{}, err
	}

	pollStats := gen.StatsData{
		Amount:  uint32(stats.Stats.Amount),
		Above70: uint32(stats.Stats.Above70),
	}
	rates := make([]uint32, 0)
	for i := range len(stats.Stats.Rates) {
		rates = append(rates, uint32(stats.Stats.Rates[i]))
	}
	pollStats.Rates = rates

	return &gen.GetPollStatsResponse{
		PollTitle: stats.PollTitle,
		Stats:     &pollStats,
	}, nil
}
