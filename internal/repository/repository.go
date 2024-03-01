package repository

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type Users interface {
	CreateUser(user models.User) (uint, error)
	GetUser(login string) (models.User, error)
}

type Sessions interface {
	CreateSession(userID uint) string
	SessionExists(sID string) bool
	DeleteSession(sID string) error
}

type Products interface {
	GetProducts(lastID, limit int) ([]models.Product, error)
}
