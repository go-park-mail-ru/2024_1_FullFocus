package dao

import "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/auth/models"

type UserTable struct {
	ID           uint   `db:"id"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
}

func ConvertTableToUser(t UserTable) models.User {
	return models.User{
		ID:           t.ID,
		Email:        t.Email,
		PasswordHash: t.PasswordHash,
	}
}

func ConvertUserToTable(m models.User) UserTable {
	return UserTable{
		ID:           m.ID,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
	}
}
