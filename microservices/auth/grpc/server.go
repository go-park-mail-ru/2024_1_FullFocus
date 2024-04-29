package authgrpc

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	authv1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/auth"
)

type Auth interface {
	Login(ctx context.Context, login string, password string) (string, error)
	Signup(ctx context.Context, login string, password string) (uint, string, error)
	GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
	Logout(ctx context.Context, sID string) error
	IsLoggedIn(ctx context.Context, sID string) bool
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

// TODO: validation on microservice side

func (s *serverAPI) Login(ctx context.Context, r *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	sID, err := s.authUsecase.Login(ctx, r.GetLogin(), r.GetPassword())
	if err != nil {
		return &authv1.LoginResponse{}, err
	}
	return &authv1.LoginResponse{
		SessionID: sID,
	}, nil
}

func (s *serverAPI) Signup(ctx context.Context, r *authv1.SignupRequest) (*authv1.SignupResponse, error) {
	uID, sID, err := s.authUsecase.Signup(ctx, r.GetLogin(), r.GetPassword())
	if err != nil {
		return &authv1.SignupResponse{}, err
	}
	return &authv1.SignupResponse{
		UserID:    uint32(uID),
		SessionID: sID,
	}, nil
}

func (s *serverAPI) Logout(ctx context.Context, r *authv1.LogoutRequest) (*empty.Empty, error) {
	return nil, s.authUsecase.Logout(ctx, r.GetSessionID())
}

func (s *serverAPI) GetUserIDBySessionID(ctx context.Context, r *authv1.GetUserIDRequest) (*authv1.GetUserIDResponse, error) {
	uID, err := s.authUsecase.GetUserIDBySessionID(ctx, r.GetSessionID())
	if err != nil {
		return &authv1.GetUserIDResponse{}, err
	}
	return &authv1.GetUserIDResponse{
		UserID: uint32(uID),
	}, nil
}

func (s *serverAPI) CheckAuth(ctx context.Context, r *authv1.CheckAuthRequest) (*authv1.CheckAuthResponse, error) {
	return &authv1.CheckAuthResponse{
		IsLoggedIn: s.authUsecase.IsLoggedIn(ctx, r.GetSessionID()),
	}, nil
}
