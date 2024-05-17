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

	promotionv1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/promotion"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type Client struct {
	api promotionv1.PromotionClient
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
		return nil, fmt.Errorf("promotion client create error: %w", err)
	}
	c := &Client{
		api: promotionv1.NewPromotionClient(conn),
	}
	return c, nil
}

func (c *Client) CreatePromoProductInfo(ctx context.Context, input models.PromoData) error {
	_, err := c.api.AddPromoProductInfo(ctx, &promotionv1.AddPromoProductRequest{
		ProductID:    uint32(input.ProductID),
		BenefitType:  input.BenefitType,
		BenefitValue: uint32(input.BenefitValue),
	})
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.OK:
		return nil
	case codes.AlreadyExists:
		return models.ErrPromoProductAlreadyExists
	case codes.NotFound:
		return models.ErrNoProduct
	case codes.InvalidArgument:
		return helper.NewValidationError(st.Proto().GetMessage(), st.Proto().GetMessage())
	default:
		return models.ErrInternal
	}
}

func (c *Client) GetPromoProductsInfo(ctx context.Context, amount uint) ([]models.PromoData, error) {
	promoResp, err := c.api.GetPromoProductsInfo(ctx, &promotionv1.GetPromoProductsRequest{
		Amount: uint32(amount),
	})
	st, ok := status.FromError(err)
	if !ok {
		return nil, err
	}
	switch st.Code() {
	case codes.OK:
		promoData := make([]models.PromoData, 0, len(promoResp.GetPromoProductsInfo()))
		for _, promo := range promoResp.GetPromoProductsInfo() {
			promoData = append(promoData, models.PromoData{
				ProductID:    uint(promo.GetProductID()),
				BenefitType:  promo.GetBenefitType(),
				BenefitValue: uint(promo.GetBenefitValue()),
			})
		}
		return promoData, nil
	case codes.NotFound:
		return nil, models.ErrNoProduct
	default:
		return nil, models.ErrInternal
	}
}

func (c *Client) DeletePromoProductInfo(ctx context.Context, ID uint) error {
	_, err := c.api.DeletePromoProductInfo(ctx, &promotionv1.DeletePromoProductRequest{
		ProductID: uint32(ID),
	})
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.OK:
		return nil
	case codes.NotFound:
		return models.ErrNoProduct
	default:
		return models.ErrInternal
	}
}
