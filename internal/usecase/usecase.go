package usecase

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type Auth interface {
	Login(login string, password string) (string, error)
	Signup(login string, password string) (string, string, error)
	Logout(sID string) error
	IsLoggedIn(isID string) bool
}

type Products interface {
	GetProducts(lastID, limit int) ([]models.Product, error)
}
