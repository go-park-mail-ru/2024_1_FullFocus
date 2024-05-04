package csatgrpc

import (
	"context"
	"errors"

	csatv1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/csat"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/models"
	commonError "github.com/go-park-mail-ru/2024_1_FullFocus/pkg/error"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CSAT interface {
	GetAllPolls(context.Context, uint) ([]models.Poll, error)
	CreatePollRate(context.Context, models.CreatePollRate) error
	GetPollStats(context.Context, uint) (models.PollStats, error)
}

type serverAPI struct {
	csatv1.UnimplementedCSATServer
	usecase CSAT
}

func Register(gRPCServer *grpc.Server, uc CSAT) {
	csatv1.RegisterCSATServer(gRPCServer, &serverAPI{
		usecase: uc,
	})
}

func (s *serverAPI) GetPolls(ctx context.Context, r *csatv1.GetPollsRequest) (*csatv1.GetPollsResponse, error) {
	polls, err := s.usecase.GetAllPolls(ctx, uint(r.GetProfileID()))
	if err != nil {
		if errors.Is(err, models.ErrNoPolls) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, commonError.ErrInternal.Error())
	}
	pollsResp := csatv1.GetPollsResponse{
		Polls: make([]*csatv1.Poll, 0),
	}
	for _, poll := range polls {
		pollsResp.Polls = append(pollsResp.Polls, &csatv1.Poll{
			Id:    uint32(poll.ID),
			Title: poll.Title,
			Voted: poll.Voted,
		})
	}
	return &pollsResp, status.Error(codes.OK, "")
}

func (s *serverAPI) CreatePollRate(ctx context.Context, r *csatv1.CreatePollRateRequest) (*empty.Empty, error) {
	input := models.CreatePollRate{
		ProfileID: uint(r.GetProfileID()),
		PollID:    uint(r.GetPollID()),
		Rate:      uint(r.GetRate()),
	}
	if err := s.usecase.CreatePollRate(ctx, input); err != nil {
		switch {
		case errors.Is(err, models.ErrNoPolls):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, models.ErrPollAlreadyRated):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, commonError.ErrInternal.Error())
		}
	}
	return nil, status.Error(codes.OK, "")
}

func (s *serverAPI) GetPollStats(ctx context.Context, r *csatv1.GetPollStatsRequest) (*csatv1.GetPollStatsResponse, error) {
	stats, err := s.usecase.GetPollStats(ctx, uint(r.GetPollID()))
	if err != nil {
		if errors.Is(err, models.ErrNoPolls) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, commonError.ErrInternal.Error())
	}
	rates := make([]uint32, 0)
	for i := range len(stats.Stats.Rates) {
		rates = append(rates, uint32(stats.Stats.Rates[i]))
	}
	return &csatv1.GetPollStatsResponse{
		PollName: stats.PollTitle,
		Amount:   uint32(stats.Stats.Amount),
		Above70:  uint32(stats.Stats.Above70),
		Rates:    rates,
	}, status.Error(codes.OK, "")
}
