package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/delivery/gen"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/models"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type CSATs interface {
	GetAllPolls(context.Context) ([]models.Poll, error)
	CreatePollRate(context.Context, models.CreatePollRate) error
}

type CSATHandler struct {
	csatUsecase CSATs
	gen.CSATServer
}

func NewCSATHandler(u CSATs) *CSATHandler {
	return &CSATHandler{
		csatUsecase: u,
	}
}

func (h *CSATHandler) GetPolls(ctx context.Context, r *empty.Empty) (*gen.GetPollsResponse, error) {
	polls, err := h.csatUsecase.GetAllPolls(ctx)
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
		})
	}
	return &payload, nil
}

func (h *CSATHandler) CreatePollRate(ctx context.Context, in *gen.CreatePollRateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	input := models.CreatePollRate{
		ProfileID: uint(in.ProfileID),
		PollID:    uint(in.PollID),
		Rate:      uint(in.Rate),
	}

	if err := h.csatUsecase.CreatePollRate(ctx, input); err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
