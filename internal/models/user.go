package models

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password"`
}
