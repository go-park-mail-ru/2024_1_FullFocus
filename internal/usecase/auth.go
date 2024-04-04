package usecase

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"

	"golang.org/x/crypto/argon2"

	"strconv"
)

const (
	_minLoginLength    = 4
	_maxLoginLength    = 32
	_minPasswordLength = 8
	_maxPasswordLength = 32
	_countBytes        = 8
	_countMemory       = 65536
	_countTreads       = 4
	_countKeyLen       = 32
)

func PasswordArgon2(plainPassword []byte, salt []byte) []byte {
	return argon2.IDKey(plainPassword, salt, 1, _countMemory, _countTreads, _countKeyLen)
}

type AuthUsecase struct {
	userRepo    repository.Users
	sessionRepo repository.Sessions
}

func NewAuthUsecase(ur repository.Users, sr repository.Sessions) *AuthUsecase {
	return &AuthUsecase{
		userRepo:    ur,
		sessionRepo: sr,
	}
}

func (u *AuthUsecase) Login(ctx context.Context, login string, password string) (string, error) {
	err := helper.ValidateField(login, _minLoginLength, _maxLoginLength)
	if err != nil {
		return "", models.NewValidationError("invalid login input",
			"Логин должен содержать от 4 до 32 букв английского алфавита или цифр")
	}
	err = helper.ValidateField(password, _minPasswordLength, _maxPasswordLength)
	if err != nil {
		return "", models.NewValidationError("invalid password input",
			"Пароль должен содержать от 8 до 32 букв английского алфавита или цифр")
	}

	user, err := u.userRepo.GetUser(ctx, login)
	if err != nil {
		return "", models.ErrNoUser
	}

	salt := ([]byte(user.Password))[0:8]
	passwordHash := PasswordArgon2([]byte(password), salt)
	saltWithPasswordHash := string(salt) + string(passwordHash)

	if saltWithPasswordHash != user.Password {
		return "", models.ErrWrongPassword
	}

	return u.sessionRepo.CreateSession(ctx, user.ID), nil
}

func (u *AuthUsecase) Signup(ctx context.Context, login string, password string) (string, string, error) {
	err := helper.ValidateField(login, _minLoginLength, _maxLoginLength)
	if err != nil {
		return "", "", models.NewValidationError("invalid login input",
			"Логин должен содержать от 4 до 32 букв английского алфавита или цифр")
	}
	err = helper.ValidateField(password, _minPasswordLength, _maxPasswordLength)
	if err != nil {
		return "", "", models.NewValidationError("invalid password input",
			"Пароль должен содержать от 8 до 32 букв английского алфавита или цифр")
	}

	salt := make([]byte, _countBytes)
	_, errSalt := rand.Read(salt)
	if errSalt != nil {
		return "", "", errors.New("err with making salt")
	}
	passwordHash := PasswordArgon2([]byte(password), salt)
	saltWithPasswordHash := string(salt) + string(passwordHash)

	user := models.User{
		Username: login,
		Password: saltWithPasswordHash,
	}

	uID, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return "", "", models.ErrUserAlreadyExists
	}
	sID := u.sessionRepo.CreateSession(ctx, uID)

	uIDHash := md5.Sum([]byte(strconv.Itoa(int(uID))))
	stringUID := hex.EncodeToString(uIDHash[:])

	return sID, stringUID, nil
}

func (u *AuthUsecase) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	return u.sessionRepo.GetUserIDBySessionID(ctx, sID)
}

func (u *AuthUsecase) Logout(ctx context.Context, sID string) error {
	return u.sessionRepo.DeleteSession(ctx, sID)
}

func (u *AuthUsecase) IsLoggedIn(ctx context.Context, sID string) bool {
	return u.sessionRepo.SessionExists(ctx, sID)
}
