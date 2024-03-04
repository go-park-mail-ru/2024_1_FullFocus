package usecase

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/pkg/errors"
)

var ErrWrongPassword = errors.New("wrong password")

type AuthUsecase struct {
	userRepo    repository.Users
	sessionRepo repository.Sessions
}

// Auth godoc
// @Tags AuthFunc
// @Summary Make new auth rep
// @Param repUsers body repository.Users true "UserID"
// @Param repSessions body repository.Sessions true "UserID"
// @Success 200 {object} AuthUsecase "new AuthUsecase"
// @Router /NewAuthUsecase [post]
func NewAuthUsecase(ur repository.Users, sr repository.Sessions) *AuthUsecase {
	return &AuthUsecase{
		userRepo:    ur,
		sessionRepo: sr,
	}
}

// Auth godoc
// @Tags AuthFunc
// @Summary Make new session
// @Success 200 {object} string
// @Success 400 {object} error
// @Param login body string true "Status code"
// @Param password body string true "Status code"
// @Router /LoginFunc [post]
func (u *AuthUsecase) Login(login string, password string) (string, error) {
	user, err := u.userRepo.GetUser(login)
	if err != nil {
		return "", err
	}
	if password != user.Password {
		return "", ErrWrongPassword
	}
	return u.sessionRepo.CreateSession(user.ID), nil
}

// Auth godoc
// @Tags AuthFunc
// @Summary Make new user and his first session
// @Success 200 {object} []string
// @Success 400 {object} error
// @Param login body string true "Status code"
// @Param password body string true "Status code"
// @Router /SignupFunc [post]
func (u *AuthUsecase) Signup(login string, password string) (string, string, error) {
	user := models.User{
		Username: login,
		Password: password,
	}
	uID, err := u.userRepo.CreateUser(user)
	if err != nil {
		return "", "", err
	}
	sID := u.sessionRepo.CreateSession(uID)

	uIDHash := md5.Sum([]byte(strconv.Itoa(int(uID))))
	stringUID := hex.EncodeToString(uIDHash[:])
	log.Printf("User ID hash: %s", stringUID)
	return sID, stringUID, nil
}

func (u *AuthUsecase) Logout(sID string) error {
	return u.sessionRepo.DeleteSession(sID)
}

func (u *AuthUsecase) IsLoggedIn(sID string) bool {
	return u.sessionRepo.SessionExists(sID)
}
