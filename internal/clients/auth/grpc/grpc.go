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

	authv1 "github.com/go-park-mail-ru/2024_1_FullFocus/gen/auth"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type Client struct {
	api authv1.AuthClient
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
		return nil, fmt.Errorf("client create error: %w", err)
	}
	c := &Client{
		api: authv1.NewAuthClient(conn),
	}
	return c, nil
}

func (c *Client) Login(ctx context.Context, login string, password string) (string, error) {
	res, err := c.api.Login(ctx, &authv1.LoginRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return "", err
	}
	return res.GetSessionID(), nil
}

func (c *Client) Signup(ctx context.Context, login string, password string) (uint, string, error) {
	res, err := c.api.Signup(ctx, &authv1.SignupRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return 0, "", err
	}
	return uint(res.GetUserID()), res.SessionID, nil
}

func (c *Client) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	res, err := c.api.GetUserIDBySessionID(ctx, &authv1.GetUserIDRequest{
		SessionID: sID,
	})
	if err != nil {
		return 0, err
	}
	return uint(res.GetUserID()), nil
}

func (c *Client) Logout(ctx context.Context, sID string) error {
	_, err := c.api.Logout(ctx, &authv1.LogoutRequest{
		SessionID: sID,
	})
	return err
}

func (c *Client) CheckAuth(ctx context.Context, sID string) bool {
	res, err := c.api.CheckAuth(ctx, &authv1.CheckAuthRequest{
		SessionID: sID,
	})
	if err != nil {
		return false
	}
	return res.GetIsLoggedIn()
}

func (c *Client) UpdatePassword(ctx context.Context, userID uint, password string, newPassword string) error {
	_, err := c.api.UpdatePassword(ctx, &authv1.UpdatePasswordRequest{
		UserID:      uint32(userID),
		Password:    password,
		NewPassword: newPassword,
	})
	return err
}
