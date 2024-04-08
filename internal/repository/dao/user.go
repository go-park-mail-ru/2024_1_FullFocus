package dao

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type UserTable struct {
	ID           uint   `db:"id"`
	Login        string `db:"user_login"`
	PasswordHash string `db:"password_hash"`
}

func ConvertTableToUser(t UserTable) models.User {
	return models.User{
		ID:           t.ID,
		Username:     t.Login,
		PasswordHash: t.PasswordHash,
	}
}

func ConvertUserToTable(m models.User) UserTable {
	return UserTable{
		ID:           m.ID,
		Login:        m.Username,
		PasswordHash: m.PasswordHash,
	}
}
