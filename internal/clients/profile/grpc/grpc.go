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

	profilev1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/profile"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type Client struct {
	api profilev1.ProfileClient
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
		return nil, fmt.Errorf("profile client create error: %w", err)
	}
	c := &Client{
		api: profilev1.NewProfileClient(conn),
	}
	return c, nil
}

func (c *Client) CreateProfile(ctx context.Context, pID uint) error {
	_, err := c.api.CreateProfile(ctx, &profilev1.CreateProfileRequest{
		ProfileID: uint32(pID),
	})
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.OK:
		return nil
	case codes.AlreadyExists:
		return models.ErrProfileAlreadyExists
	default:
		return st.Err()
	}
}

func (c *Client) GetProfileByID(ctx context.Context, pID uint) (models.Profile, error) {
	res, err := c.api.GetProfileByID(ctx, &profilev1.GetProfileByIDRequest{
		ProfileID: uint32(pID),
	})
	st, ok := status.FromError(err)
	if !ok {
		return models.Profile{}, err
	}
	switch st.Code() {
	case codes.OK:
		profile := models.Profile{
			ID:         pID,
			FullName:   res.GetName(),
			Address:    res.GetAddress(),
			PhoneNum:   res.GetPhoneNum(),
			Gender:     uint(res.GetGender()),
			AvatarName: res.GetAvatarName(),
		}
		return profile, nil
	case codes.NotFound:
		return models.Profile{}, models.ErrNoProfile
	default:
		return models.Profile{}, st.Err()
	}
}

func (c *Client) GetProfileNamesByIDs(ctx context.Context, pIDs []uint) ([]string, error) {
	var pIDs32 []uint32
	for _, pID := range pIDs {
		pIDs32 = append(pIDs32, uint32(pID))
	}
	res, err := c.api.GetProfileNamesByIDs(ctx, &profilev1.GetProfileNamesByIDsRequest{
		ProfileIDs: pIDs32,
	})
	st, ok := status.FromError(err)
	if !ok {
		return nil, err
	}
	switch st.Code() {
	case codes.OK:
		return res.GetNames(), nil
	case codes.NotFound:
		return nil, models.ErrNoProfile
	default:
		return nil, st.Err()
	}
}

func (c *Client) GetProfileMetaInfo(ctx context.Context, pID uint) (models.ProfileMetaInfo, error) {
	res, err := c.api.GetProfileMetaInfo(ctx, &profilev1.GetProfileMetaInfoRequest{
		ProfileID: uint32(pID),
	})
	st, ok := status.FromError(err)
	if !ok {
		return models.ProfileMetaInfo{}, err
	}
	switch st.Code() {
	case codes.OK:
		info := models.ProfileMetaInfo{
			FullName:   res.GetProfileName(),
			AvatarName: res.GetAvatarName(),
		}
		return info, nil
	case codes.NotFound:
		return models.ProfileMetaInfo{}, models.ErrNoProfile
	default:
		return models.ProfileMetaInfo{}, st.Err()
	}
}

func (c *Client) GetAvatarByID(ctx context.Context, pID uint) (string, error) {
	res, err := c.api.GetAvatarByID(ctx, &profilev1.GetAvatarByIDRequest{
		ProfileID: uint32(pID),
	})
	st, ok := status.FromError(err)
	if !ok {
		return "", err
	}
	switch st.Code() {
	case codes.OK:
		return res.GetAvatarName(), nil
	case codes.NotFound:
		return "", models.ErrNoProfile
	default:
		return "", st.Err()
	}
}

func (c *Client) UpdateAvatarByProfileID(ctx context.Context, pID uint, avatarName string) (string, error) {
	res, err := c.api.UpdateAvatarByProfileID(ctx, &profilev1.UpdateAvatarByProfileIDRequest{
		ProfileID:  uint32(pID),
		AvatarName: avatarName,
	})
	st, ok := status.FromError(err)
	if !ok {
		return "", err
	}
	switch st.Code() {
	case codes.OK:
		return res.GetPrevAvatarName(), nil
	default:
		return "", st.Err()
	}
}

func (c *Client) UpdateProfile(ctx context.Context, pID uint, newProfile models.ProfileUpdateInput) error {
	_, err := c.api.UpdateProfile(ctx, &profilev1.UpdateProfileRequest{
		ProfileID: uint32(pID),
		Name:      newProfile.FullName,
		Address:   newProfile.Address,
		PhoneNum:  newProfile.PhoneNum,
		Gender:    uint32(newProfile.Gender),
	})
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.OK:
		return nil
	case codes.NotFound:
		return models.ErrNoProfile
	case codes.InvalidArgument:
		return models.ErrInvalidField
	default:
		return st.Err()
	}
}

func (c *Client) DeleteAvatarByProfileID(ctx context.Context, pID uint) (string, error) {
	res, err := c.api.DeleteAvatarByProfileID(ctx, &profilev1.DeleteAvatarByProfileIDRequest{
		ProfileID: uint32(pID),
	})
	st, ok := status.FromError(err)
	if !ok {
		return "", err
	}
	switch st.Code() {
	case codes.OK:
		return res.GetPrevAvatarName(), nil
	case codes.NotFound:
		return "", models.ErrNoProfile
	default:
		return "", st.Err()
	}
}

func (c *Client) GetProfileNamesAvatarsByIDs(ctx context.Context, pIDs []uint) ([]models.ProfileNameAvatar, error) {
	var pIDs32 []uint32
	for _, pID := range pIDs {
		pIDs32 = append(pIDs32, uint32(pID))
	}
	resp, err := c.api.GetProfileNamesAvatarsByIDs(ctx, &profilev1.GetProfileNamesAvatarsRequest{
		ProfileIDs: pIDs32,
	})
	st, ok := status.FromError(err)
	if !ok {
		return nil, err
	}
	switch st.Code() {
	case codes.OK:
		data := make([]models.ProfileNameAvatar, 0)
		for _, r := range resp.GetData() {
			data = append(data, models.ProfileNameAvatar{
				FullName:   r.GetName(),
				AvatarName: r.GetAvatar(),
			})
		}
		return data, nil
	case codes.NotFound:
		return nil, models.ErrNoProfile
	default:
		return nil, st.Err()
	}
}
