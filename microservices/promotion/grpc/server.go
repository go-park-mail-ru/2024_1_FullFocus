package promotiongrpc

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	promotionv1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/promotion"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/models"
	commonError "github.com/go-park-mail-ru/2024_1_FullFocus/pkg/error"
)

type Promotion interface {
	CreatePromoProductInfo(ctx context.Context, input models.PromoData) error
	GetPromoProductsInfoByIDs(ctx context.Context, prIDs []uint) ([]models.PromoData, error)
	DeletePromoProductInfo(ctx context.Context, pID uint) error
}

type serverAPI struct {
	promotionv1.UnimplementedPromotionServer
	usecase Promotion
}

func Register(gRPCServer *grpc.Server, uc Promotion) {
	promotionv1.RegisterPromotionServer(gRPCServer, &serverAPI{
		usecase: uc,
	})
}

func (s *serverAPI) AddPromoProductInfo(ctx context.Context, r *promotionv1.AddPromoProductRequest) (*empty.Empty, error) {
	if err := s.usecase.CreatePromoProductInfo(ctx, models.PromoData{
		ProductID:    uint(r.GetProductID()),
		BenefitType:  r.GetBenefitType(),
		BenefitValue: uint(r.GetBenefitValue()),
	}); err != nil {
		switch {
		case errors.Is(err, models.ErrProductNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, models.ErrPromoProductAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, commonError.ErrInternal.Error())
		}
	}
	return nil, status.Error(codes.OK, "")
}

func (s *serverAPI) GetPromoProductsInfoByIDs(ctx context.Context, r *promotionv1.GetPromoProductsRequest) (*promotionv1.GetPromoProductsResponse, error) {
	uint32PrIDs := r.GetProductIDs()
	prIDs := make([]uint, 0, len(uint32PrIDs))
	for _, id := range uint32PrIDs {
		prIDs = append(prIDs, uint(id))
	}
	promoResp, err := s.usecase.GetPromoProductsInfoByIDs(ctx, prIDs)
	if err != nil {
		if errors.Is(err, models.ErrProductNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	promoProductsInfo := make([]*promotionv1.PromoData, 0, len(promoResp))
	for _, promoInfo := range promoResp {
		promoProductsInfo = append(promoProductsInfo, &promotionv1.PromoData{
			ProductID:    uint32(promoInfo.ProductID),
			BenefitType:  promoInfo.BenefitType,
			BenefitValue: uint32(promoInfo.BenefitValue),
		})
	}
	return &promotionv1.GetPromoProductsResponse{
		PromoProductsInfo: promoProductsInfo,
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) DeletePromoProductInfo(ctx context.Context, r *promotionv1.DeletePromoProductRequest) (*empty.Empty, error) {
	if err := s.usecase.DeletePromoProductInfo(ctx, uint(r.GetProductID())); err != nil {
		if errors.Is(err, models.ErrProductNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, commonError.ErrInternal.Error())
	}
	return nil, status.Error(codes.OK, "")
}
