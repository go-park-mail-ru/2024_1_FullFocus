package authgrpc

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authv1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/auth"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/models"
)

type Auth interface {
	Login(ctx context.Context, login string, password string) (string, error)
	Signup(ctx context.Context, login string, password string) (uint, string, error)
	GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
	GetUserLoginByID(ctx context.Context, uID uint) (string, error)
	Logout(ctx context.Context, sID string) error
	IsLoggedIn(ctx context.Context, sID string) bool
	UpdatePassword(ctx context.Context, userID uint, password string, newPassword string) error
}

type serverAPI struct {
	authv1.UnimplementedAuthServer
	authUsecase Auth
}

func Register(gRPCServer *grpc.Server, uc Auth) {
	authv1.RegisterAuthServer(gRPCServer, &serverAPI{
		authUsecase: uc,
	})
}

func (s *serverAPI) Login(ctx context.Context, r *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	sID, err := s.authUsecase.Login(ctx, r.GetLogin(), r.GetPassword())
	if err != nil {
		if errors.Is(err, models.ErrInvalidInput) {
			return &authv1.LoginResponse{}, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, models.ErrUserNotFound) {
			return &authv1.LoginResponse{}, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, models.ErrWrongPassword) {
			return &authv1.LoginResponse{}, status.Error(codes.PermissionDenied, err.Error())
		}
	}
	return &authv1.LoginResponse{
		SessionID: sID,
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) Signup(ctx context.Context, r *authv1.SignupRequest) (*authv1.SignupResponse, error) {
	uID, sID, err := s.authUsecase.Signup(ctx, r.GetLogin(), r.GetPassword())
	if err != nil {
		if errors.Is(err, models.ErrInvalidInput) {
			return &authv1.SignupResponse{}, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, models.ErrUnableToHash) {
			return &authv1.SignupResponse{}, status.Error(codes.Internal, err.Error())
		} else if errors.Is(err, models.ErrUserAlreadyExists) {
			return &authv1.SignupResponse{}, status.Error(codes.AlreadyExists, err.Error())
		}
	}
	return &authv1.SignupResponse{
		UserID:    uint32(uID),
		SessionID: sID,
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) Logout(ctx context.Context, r *authv1.LogoutRequest) (*empty.Empty, error) {
	err := s.authUsecase.Logout(ctx, r.GetSessionID())
	if err != nil {
		return &empty.Empty{}, status.Error(codes.PermissionDenied, err.Error())
	}
	return &empty.Empty{}, status.Error(codes.OK, "")
}

func (s *serverAPI) GetUserIDBySessionID(ctx context.Context, r *authv1.GetUserIDRequest) (*authv1.GetUserIDResponse, error) {
	uID, err := s.authUsecase.GetUserIDBySessionID(ctx, r.GetSessionID())
	if err != nil {
		return &authv1.GetUserIDResponse{}, status.Error(codes.PermissionDenied, err.Error())
	}
	return &authv1.GetUserIDResponse{
		UserID: uint32(uID),
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) GetUserLoginByID(ctx context.Context, r *authv1.GetUserLoginByIDRequest) (*authv1.GetUserLoginByIDResponse, error) {
	login, err := s.authUsecase.GetUserLoginByID(ctx, uint(r.GetUserID()))
	if err != nil {
		return &authv1.GetUserLoginByIDResponse{}, status.Error(codes.NotFound, err.Error())
	}
	return &authv1.GetUserLoginByIDResponse{
		Login: login,
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) CheckAuth(ctx context.Context, r *authv1.CheckAuthRequest) (*authv1.CheckAuthResponse, error) {
	return &authv1.CheckAuthResponse{
		IsLoggedIn: s.authUsecase.IsLoggedIn(ctx, r.GetSessionID()),
	}, status.Error(codes.OK, "")
}

func (s *serverAPI) UpdatePassword(ctx context.Context, r *authv1.UpdatePasswordRequest) (*empty.Empty, error) {
	err := s.authUsecase.UpdatePassword(ctx, uint(r.GetUserID()), r.GetPassword(), r.GetNewPassword())
	if err != nil {
		if errors.Is(err, models.ErrInvalidInput) {
			return &empty.Empty{}, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, models.ErrUserNotFound) {
			return &empty.Empty{}, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, models.ErrWrongPassword) {
			return &empty.Empty{}, status.Error(codes.PermissionDenied, err.Error())
		} else if errors.Is(err, models.ErrUnableToHash) {
			return &empty.Empty{}, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, status.Error(codes.OK, "")
}
