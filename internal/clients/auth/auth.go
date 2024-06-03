package auth

import "context"

type AuthClient interface {
	CheckAuth(ctx context.Context, sID string) bool
	GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
	GetUserLoginByID(ctx context.Context, uID uint) (string, error)
	Login(ctx context.Context, login string, password string) (string, error)
	Logout(ctx context.Context, sID string) error
	Signup(ctx context.Context, login string, password string) (uint, string, error)
	UpdatePassword(ctx context.Context, userID uint, password string, newPassword string) error
}
