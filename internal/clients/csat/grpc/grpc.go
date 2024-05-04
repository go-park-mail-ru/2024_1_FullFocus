package grpc

import (
	"context"
	"fmt"
	"log/slog"

	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	csatv1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/csat"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type Client struct {
	api csatv1.CSATClient
}

func New(ctx context.Context, log *slog.Logger, cfg config.ClientConfig) (*Client, error) {
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.Aborted, codes.DeadlineExceeded, codes.NotFound),
		grpcretry.WithMax(uint(cfg.Retries)),
		grpcretry.WithPerRetryTimeout(cfg.RetryTimeout),
	}
	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}
	conn, err := grpc.DialContext(ctx,
		cfg.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(logger.InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("CSAT client create error: %w", err)
	}
	c := &Client{
		api: csatv1.NewCSATClient(conn),
	}
	return c, nil
}

func (c *Client) GetAllPolls(ctx context.Context, profileID uint) ([]models.Poll, error) {
	pollsResp, err := c.api.GetPolls(ctx, &csatv1.GetPollsRequest{
		ProfileID: uint32(profileID),
	})
	st, ok := status.FromError(err)
	if !ok {
		return nil, err
	}
	switch st.Code() {
	case codes.OK:
		polls := make([]models.Poll, 0)
		for _, p := range pollsResp.GetPolls() {
			polls = append(polls, models.Poll{
				ID:    uint(p.GetId()),
				Title: p.GetTitle(),
				Voted: p.GetVoted(),
			})
		}
		return polls, nil
	case codes.NotFound:
		return nil, models.ErrNoPolls
	default:
		return nil, st.Err()
	}
}

func (c *Client) CreatePollRate(ctx context.Context, input models.CreatePollRateInput) error {
	_, err := c.api.CreatePollRate(ctx, &csatv1.CreatePollRateRequest{
		ProfileID: uint32(input.ProfileID),
		PollID:    uint32(input.PollID),
		Rate:      uint32(input.Rate),
	})
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.OK:
		return nil
	case codes.NotFound:
		return models.ErrNoPolls
	case codes.AlreadyExists:
		return models.ErrPollAlreadyRated
	default:
		return st.Err()
	}
}

func (c *Client) GetPollStats(ctx context.Context, pollID uint) (models.PollStats, error) {
	ratesResp, err := c.api.GetPollStats(ctx, &csatv1.GetPollStatsRequest{
		PollID: uint32(pollID),
	})
	st, ok := status.FromError(err)
	if !ok {
		return models.PollStats{}, err
	}
	switch st.Code() {
	case codes.OK:
		return models.PollStats{
			Title:   ratesResp.GetPollName(),
			Rates:   ratesResp.GetRates(),
			Amount:  ratesResp.GetAmount(),
			Above70: ratesResp.GetAbove70(),
		}, nil
	case codes.NotFound:
		return models.PollStats{}, models.ErrNoPolls
	case codes.AlreadyExists:
		return models.PollStats{}, models.ErrPollAlreadyRated
	default:
		return models.PollStats{}, st.Err()
	}
}
