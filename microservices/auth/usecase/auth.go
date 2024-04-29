package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
)

type auth interface {
	CreateUser(ctx context.Context, user models.User) (uint, error)
	GetUser(ctx context.Context, login string) (models.User, error)
	CreateSession(ctx context.Context, userID uint) string
	GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
	SessionExists(ctx context.Context, sID string) bool
	DeleteSession(ctx context.Context, sID string) error
}

type Auth struct {
	repo auth
}

func NewAuthUsecase(a auth) *Auth {
	return &Auth{
		repo: a,
	}
}

func (u *Auth) Login(ctx context.Context, login string, password string) (string, error) {
	user, err := u.repo.GetUser(ctx, login)
	if err != nil {
		return "", models.ErrNoUser
	}
	if err = helper.CheckPassword(password, user.PasswordHash); err != nil {
		return "", models.ErrWrongPassword
	}
	return u.repo.CreateSession(ctx, user.ID), nil
}

func (u *Auth) Signup(ctx context.Context, login string, password string) (uint, string, error) {
	passwordHash, err := helper.HashPassword(password)
	if err != nil {
		return 0, "", err
	}
	user := models.User{
		Username:     login,
		PasswordHash: passwordHash,
	}
	uID, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return 0, "", models.ErrUserAlreadyExists
	}
	sID := u.repo.CreateSession(ctx, uID)
	return uID, sID, nil
}

func (u *Auth) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	return u.repo.GetUserIDBySessionID(ctx, sID)
}

func (u *Auth) Logout(ctx context.Context, sID string) error {
	return u.repo.DeleteSession(ctx, sID)
}

func (u *Auth) IsLoggedIn(ctx context.Context, sID string) bool {
	return u.repo.SessionExists(ctx, sID)
}
