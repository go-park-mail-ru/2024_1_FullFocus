package usecase

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/pkg/errors"
	"log"
	"strconv"
)

var ErrWrongPassword = errors.New("wrong password")

type AuthUsecase struct {
	userRepo    repository.Users
	sessionRepo repository.Sessions
}

func (u *AuthUsecase) Login(login string, password string) (string, error) {
	user, err := u.userRepo.GetUser(login)
	if err != nil {
		return "", err
	}
	if password != user.Password {
		return "", ErrWrongPassword
	}
	return u.sessionRepo.CreateSession(login, user.ID), nil
}

func (u *AuthUsecase) Signup(login string, password string) (string, string, error) {
	user := models.User{
		Username: login,
		Password: password,
	}
	uID, err := u.userRepo.CreateUser(user)
	if err != nil {
		return "", "", err
	}
	sID := u.sessionRepo.CreateSession(login, uID)

	uIDHash := md5.Sum([]byte(strconv.Itoa(int(uID))))
	stringUID := hex.EncodeToString(uIDHash[:])
	log.Printf("User ID hash: %s", stringUID)
	return sID, stringUID, nil
}

func (u *AuthUsecase) Logout(login string) error {
	return u.sessionRepo.DeleteSession(login)
}

func (u *AuthUsecase) IsLoggedIn(login string) bool {
	return u.sessionRepo.SessionExists(login)
}
