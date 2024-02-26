package repository

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type Users interface {
	CreateUser(user models.User) (uint, error)
	GetUser(login string) (models.User, error)
}

type Sessions interface {
	CreateSession(login string, userID uint) string
	SessionExists(login string) bool
	DeleteSession(login string) error
}
