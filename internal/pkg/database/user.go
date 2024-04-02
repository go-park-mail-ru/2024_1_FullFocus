package database

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/google/uuid"
)

type UserTable struct {
	ID           uuid.UUID `db:"id"`
	Login        string    `db:"user_login"`
	PasswordHash string    `db:"password_hash"`
}

func ConvertTableToUser(t UserTable) models.User {
	return models.User{
		ID:       t.ID,
		Username: t.Login,
	}
}

func ConvertUserToTable(m models.User) UserTable {
	return UserTable{
		ID:           m.ID,
		Login:        m.Username,
		PasswordHash: m.PasswordHash,
	}
}
