package reviewgrpc

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	reviewv1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/review"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/review/models"
	commonError "github.com/go-park-mail-ru/2024_1_FullFocus/pkg/error"
)

type Review interface {
	CreateProductReview(ctx context.Context, input models.CreateProductReviewInput) error
	GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReviewData, error)
}

type serverAPI struct {
	reviewv1.UnimplementedReviewServer
	usecase Review
}

func Register(gRPCServer *grpc.Server, uc Review) {
	reviewv1.RegisterReviewServer(gRPCServer, &serverAPI{
		usecase: uc,
	})
}

func (s *serverAPI) CreateProductReview(ctx context.Context, r *reviewv1.CreateProductReviewRequest) (*empty.Empty, error) {
	if err := s.usecase.CreateProductReview(ctx, models.CreateProductReviewInput{
		ProductID:     uint(r.GetProductID()),
		ProfileID:     uint(r.GetProfileID()),
		Rating:        uint(r.GetReviewData().GetRating()),
		Advanatages:   r.GetReviewData().GetAdvantages(),
		Disadvantages: r.GetReviewData().GetDisadvantages(),
		Comment:       r.GetReviewData().GetComment(),
	}); err != nil {
		switch {
		case errors.Is(err, models.ErrReviewAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case errors.Is(err, models.ErrNoProduct):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, commonError.ErrInvalidInput):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, commonError.ErrInternal):
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, status.Error(codes.OK, "")
}

func (s *serverAPI) GetProductReviews(ctx context.Context, r *reviewv1.GetProductReviewsRequest) (*reviewv1.GetProductReviewsResponse, error) {
	reviewsData, err := s.usecase.GetProductReviews(ctx, models.GetProductReviewsInput{
		ProductID:    uint(r.GetProductID()),
		LastReviewID: uint(r.GetLastReviewID()),
		Limit:        uint(r.GetLimit()),
	})
	if err != nil {
		if errors.Is(err, models.ErrNoReviews) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	reviewsResp := make([]*reviewv1.ProductReview, 0)
	for _, r := range reviewsData {
		reviewsResp = append(reviewsResp, &reviewv1.ProductReview{
			ReviewID:  uint32(r.ReviewID),
			ProfileID: uint32(r.ProfileID),
			CreatedAt: r.CreatedAt,
			ReviewData: &reviewv1.ProductReviewData{
				Rating:        uint32(r.Rating),
				Advantages:    r.Advanatages,
				Disadvantages: r.Disadvantages,
				Comment:       r.Comment,
			},
		})
	}
	return &reviewv1.GetProductReviewsResponse{
		Reviews: reviewsResp,
	}, nil
}
