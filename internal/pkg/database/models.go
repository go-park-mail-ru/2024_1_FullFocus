package database

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/google/uuid"
)

type UserTable struct {
	Id            uuid.UUID `db:"id"`
	Login         string    `db:"user_login"`
	Password_hash string    `db:"password_hash"`
}

func ConvertTableToUser(t UserTable) models.User {
	return models.User{
		ID:       t.Id,
		Username: t.Login,
	}
}

func ConvertUserToTable(m models.User) UserTable {
	return UserTable{
		Id:            m.ID,
		Login:         m.Username,
		Password_hash: m.Password_hash,
	}
}
