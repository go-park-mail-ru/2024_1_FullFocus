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

	reviewv1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/review"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	commonError "github.com/go-park-mail-ru/2024_1_FullFocus/pkg/error"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type Client struct {
	api reviewv1.ReviewClient
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
		return nil, fmt.Errorf("review client create error: %w", err)
	}
	c := &Client{
		api: reviewv1.NewReviewClient(conn),
	}
	return c, nil
}

func (c *Client) CreateProductReview(ctx context.Context, uID uint, input models.CreateProductReviewInput) error {
	_, err := c.api.CreateProductReview(ctx, &reviewv1.CreateProductReviewRequest{
		ProfileID: uint32(uID),
		ProductID: uint32(input.ProductID),
		ReviewData: &reviewv1.ProductReviewData{
			Rating:        uint32(input.Rating),
			Advantages:    input.Advanatages,
			Disadvantages: input.Disadvantages,
			Comment:       input.Comment,
		},
	})
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.OK:
		return nil
	case codes.AlreadyExists:
		return models.ErrReviewAlreadyExists
	case codes.NotFound:
		return models.ErrNoProduct
	default:
		return commonError.ErrInternal
	}
}

func (c *Client) GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReviewData, error) {
	reviewsResp, err := c.api.GetProductReviews(ctx, &reviewv1.GetProductReviewsRequest{
		ProductID:    uint32(input.ProductID),
		LastReviewID: uint32(input.LastReviewID),
		Limit:        uint32(input.PageSize),
	})
	st, ok := status.FromError(err)
	if !ok {
		return nil, err
	}
	switch st.Code() {
	case codes.OK:
		reviews := make([]models.ProductReviewData, 0)
		for _, r := range reviewsResp.GetReviews() {
			reviews = append(reviews, models.ProductReviewData{
				ReviewID:      uint(r.GetReviewID()),
				ProfileID:     uint(r.GetProfileID()),
				Rating:        uint(r.GetReviewData().GetRating()),
				Advanatages:   r.GetReviewData().GetAdvantages(),
				Disadvantages: r.GetReviewData().GetDisadvantages(),
				Comment:       r.GetReviewData().GetComment(),
				CreatedAt:     r.GetCreatedAt(),
			})
		}
		return reviews, nil
	case codes.NotFound:
		return nil, models.ErrNoProduct
	default:
		return nil, commonError.ErrInternal
	}
}
