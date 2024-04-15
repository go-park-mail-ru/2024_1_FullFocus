package dao

import (
	model "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type ProfileTable struct {
	ID          uint   `db:"id"`
	FullName    string `db:"full_name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
}

func ConvertTableToProfile(t ProfileTable) model.Profile {
	return model.Profile{
		ID:          t.ID,
		FullName:    t.FullName,
		Email:       t.Email,
		PhoneNumber: t.PhoneNumber,
	}
}
