package profilegrpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	profilev1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/profile"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/models"
)

//go:generate mockgen -source=server.go -destination=../usecase/mocks/usecase_mock.go
type Profile interface {
	CreateProfile(ctx context.Context, pID uint) error
	UpdateProfile(ctx context.Context, uID uint, newProfile models.ProfileUpdateInput) error
	GetProfile(ctx context.Context, uID uint) (models.Profile, error)
	GetProfileNamesByIDs(ctx context.Context, pIDs []uint) ([]string, error)
	GetProfileMetaInfo(ctx context.Context, pID uint) (models.ProfileMetaInfo, error)
	GetProfileNamesAvatarsByIDs(ctx context.Context, pIDs []uint) ([]models.ProfileNameAvatar, error)
	UpdateAvatarByProfileID(ctx context.Context, uID uint, imgSrc string) (string, error)
	GetAvatarByProfileID(ctx context.Context, uID uint) (string, error)
	DeleteAvatarByProfileID(ctx context.Context, uID uint) (string, error)
}

type serverAPI struct {
	profilev1.UnimplementedProfileServer
	usecase Profile
}

func Register(gRPCServer *grpc.Server, uc Profile) {
	profilev1.RegisterProfileServer(gRPCServer, &serverAPI{
		usecase: uc,
	})
}

func (s *serverAPI) CreateProfile(ctx context.Context, r *profilev1.CreateProfileRequest) (*empty.Empty, error) {
	pID := uint(r.GetProfileID())
	if err := s.usecase.CreateProfile(ctx, pID); err != nil {
		if errors.Is(err, models.ErrProfileAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
	}
	return nil, status.Error(codes.OK, "")
}

func (s *serverAPI) GetProfileByID(ctx context.Context, r *profilev1.GetProfileByIDRequest) (*profilev1.GetProfileByIDResponse, error) {
	profile, err := s.usecase.GetProfile(ctx, uint(r.GetProfileID()))
	if err != nil {
		if errors.Is(err, models.ErrNoProfile) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
	}
	profileResp := &profilev1.GetProfileByIDResponse{
		Name:       profile.FullName,
		Address:    profile.Address,
		PhoneNum:   profile.PhoneNumber,
		Gender:     uint32(profile.Gender),
		AvatarName: profile.AvatarName,
	}
	return profileResp, status.Error(codes.OK, "")
}

func (s *serverAPI) GetProfileNamesByIDs(ctx context.Context, r *profilev1.GetProfileNamesByIDsRequest) (*profilev1.GetProfileNamesByIDsResponse, error) {
	var profileIDs []uint
	for _, id := range r.GetProfileIDs() {
		profileIDs = append(profileIDs, uint(id))
	}
	names, err := s.usecase.GetProfileNamesByIDs(ctx, profileIDs)
	if err != nil {
		if errors.Is(err, models.ErrNoProfile) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
	}
	return &profilev1.GetProfileNamesByIDsResponse{
		Names: names,
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) GetProfileMetaInfo(ctx context.Context, r *profilev1.GetProfileMetaInfoRequest) (*profilev1.GetProfileMetaInfoResponse, error) {
	info, err := s.usecase.GetProfileMetaInfo(ctx, uint(r.GetProfileID()))
	if err != nil {
		if errors.Is(err, models.ErrNoProfile) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
	}
	return &profilev1.GetProfileMetaInfoResponse{
		ProfileName: info.FullName,
		AvatarName:  info.AvatarName,
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) GetAvatarByID(ctx context.Context, r *profilev1.GetAvatarByIDRequest) (*profilev1.GetAvatarByIDResponse, error) {
	avatar, err := s.usecase.GetAvatarByProfileID(ctx, uint(r.GetProfileID()))
	if err != nil {
		if errors.Is(err, models.ErrNoProfile) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
	}
	return &profilev1.GetAvatarByIDResponse{
		AvatarName: avatar,
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) GetProfileNamesAvatarsByIDs(ctx context.Context, r *profilev1.GetProfileNamesAvatarsRequest) (*profilev1.GetProfileNamesAvatarsResponse, error) {
	pIDs := make([]uint, 0)
	for _, id := range r.GetProfileIDs() {
		pIDs = append(pIDs, uint(id))
	}
	profiles, err := s.usecase.GetProfileNamesAvatarsByIDs(ctx, pIDs)
	if err != nil {
		if errors.Is(err, models.ErrNoProfile) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	data := make([]*profilev1.ProfileNameAvatar, 0)
	for _, p := range profiles {
		data = append(data, &profilev1.ProfileNameAvatar{
			Name:   p.FullName,
			Avatar: p.AvatarName,
		})
	}
	return &profilev1.GetProfileNamesAvatarsResponse{
		Data: data,
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) UpdateAvatarByProfileID(ctx context.Context, r *profilev1.UpdateAvatarByProfileIDRequest) (*profilev1.UpdateAvatarByProfileIDResponse, error) {
	prevAvatar, err := s.usecase.UpdateAvatarByProfileID(ctx, uint(r.GetProfileID()), r.GetAvatarName())
	if err != nil {
		if errors.Is(err, models.ErrNoProfile) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
	}
	return &profilev1.UpdateAvatarByProfileIDResponse{
		PrevAvatarName: prevAvatar,
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) UpdateProfile(ctx context.Context, r *profilev1.UpdateProfileRequest) (*empty.Empty, error) {
	newProfile := models.ProfileUpdateInput{
		FullName:    r.GetName(),
		Address:     r.GetAddress(),
		PhoneNumber: r.GetPhoneNum(),
		Gender:      uint(r.GetGender()),
	}
	if err := s.usecase.UpdateProfile(ctx, uint(r.GetProfileID()), newProfile); err != nil {
		if errors.Is(err, models.ErrNoProfile) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, models.ErrInvalidInput) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, models.ErrPhoneAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
	}
	return nil, status.Error(codes.OK, "")
}

func (s *serverAPI) DeleteAvatarByProfileID(ctx context.Context, r *profilev1.DeleteAvatarByProfileIDRequest) (*profilev1.DeleteAvatarByProfileIDResponse, error) {
	prevAvatar, err := s.usecase.DeleteAvatarByProfileID(ctx, uint(r.GetProfileID()))
	if err != nil {
		if errors.Is(err, models.ErrNoProfile) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
	}
	return &profilev1.DeleteAvatarByProfileIDResponse{
		PrevAvatarName: prevAvatar,
	}, status.Error(codes.OK, "")
}
