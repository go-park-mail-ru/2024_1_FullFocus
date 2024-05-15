package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/pkg/helper"
)

type Auth interface {
	CreateUser(ctx context.Context, user models.User) (uint, error)
	GetUser(ctx context.Context, email string) (models.User, error)
	CreateSession(ctx context.Context, userID uint) string
	GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
	GetUserEmailByUserID(ctx context.Context, uID uint) (string, error)
	SessionExists(ctx context.Context, sID string) bool
	DeleteSession(ctx context.Context, sID string) error
	UpdatePassword(ctx context.Context, userID uint, password string) error
	GetUserPassword(ctx context.Context, userID uint) (string, error)
}

type Usecase struct {
	repo Auth
}

func NewAuthUsecase(a Auth) *Usecase {
	return &Usecase{
		repo: a,
	}
}

func (u *Usecase) Login(ctx context.Context, email, password string) (string, error) {
	if email == "" || password == "" {
		return "", models.ErrInvalidInput
	}
	user, err := u.repo.GetUser(ctx, email)
	if err != nil {
		return "", err
	}
	if err = helper.CheckPassword(password, user.PasswordHash); err != nil {
		return "", models.ErrWrongPassword
	}
	return u.repo.CreateSession(ctx, user.ID), nil
}

func (u *Usecase) Signup(ctx context.Context, email, password string) (uint, string, error) {
	if email == "" || password == "" {
		return 0, "", models.ErrInvalidInput
	}
	passwordHash, err := helper.HashPassword(password)
	if err != nil {
		return 0, "", models.ErrUnableToHash
	}
	user := models.User{
		Email:        email,
		PasswordHash: passwordHash,
	}
	uID, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return 0, "", err
	}
	sID := u.repo.CreateSession(ctx, uID)
	return uID, sID, nil
}

func (u *Usecase) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	return u.repo.GetUserIDBySessionID(ctx, sID)
}

func (u *Usecase) GetUserEmailByUserID(ctx context.Context, uID uint) (string, error) {
	return u.repo.GetUserEmailByUserID(ctx, uID)
}

func (u *Usecase) Logout(ctx context.Context, sID string) error {
	return u.repo.DeleteSession(ctx, sID)
}

func (u *Usecase) IsLoggedIn(ctx context.Context, sID string) bool {
	return u.repo.SessionExists(ctx, sID)
}

func (u *Usecase) UpdatePassword(ctx context.Context, userID uint, password string, newPassword string) error {
	if password == "" || newPassword == "" {
		return models.ErrInvalidInput
	}
	prevPassword, err := u.repo.GetUserPassword(ctx, userID)
	if err != nil {
		return err
	}
	if err = helper.CheckPassword(password, prevPassword); err != nil {
		return models.ErrWrongPassword
	}
	passwordHash, err := helper.HashPassword(newPassword)
	if err != nil {
		return models.ErrUnableToHash
	}
	return u.repo.UpdatePassword(ctx, userID, passwordHash)
}
