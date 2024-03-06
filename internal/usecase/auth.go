package usecase

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

const (
	_minLoginLength    = 4
	_maxLoginLength    = 32
	_minPasswordLength = 8
	_maxPasswordLength = 32
)

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

func (u *AuthUsecase) Login(login string, password string) (string, error) {
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
	user, err := u.userRepo.GetUser(login)
	if err != nil {
		return "", models.ErrNoUser // models.NewValidationError(err.Error(), "Пользователь не найден")
	}
	if password != user.Password {
		return "", models.ErrWrongPassword // models.NewValidationError("wrong password", "Неверный пароль")
	}
	return u.sessionRepo.CreateSession(user.ID), nil
}

func (u *AuthUsecase) Signup(login string, password string) (string, string, error) {
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
	user := models.User{
		Username: login,
		Password: password,
	}
	uID, err := u.userRepo.CreateUser(user)
	if err != nil {
		return "", "", models.ErrUserAlreadyExists // models.NewValidationError(err.Error(), "Пользователь с таким логином уже существует")
	}
	sID := u.sessionRepo.CreateSession(uID)

	uIDHash := md5.Sum([]byte(strconv.Itoa(int(uID))))
	stringUID := hex.EncodeToString(uIDHash[:])
	log.Printf("User ID hash: %s", stringUID)
	return sID, stringUID, nil
}

func (u *AuthUsecase) GetUserIDBySessionID(sID string) (uint, error) {
	return u.sessionRepo.GetUserIDBySessionID(sID)
}

func (u *AuthUsecase) Logout(sID string) error {
	return u.sessionRepo.DeleteSession(sID)
}

func (u *AuthUsecase) IsLoggedIn(sID string) bool {
	return u.sessionRepo.SessionExists(sID)
}
