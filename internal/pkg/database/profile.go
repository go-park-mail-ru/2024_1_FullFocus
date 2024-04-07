package database

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type ProfileTable struct {
	ID          uint   `db:"id"`
	Email       string `db:"email"`
	FullName    string `db:"full_name"`
	PhoneNumber string `db:"phone_number"`
	ImgSrc      string `db:"imgsrc"`
}

type ProfileUpdateTable struct {
	ID          uint   `db:"id"`
	Email       string `db:"email"`
	FullName    string `db:"full_name"`
	PhoneNumber string `db:"phone_number"`
	ImgSrc      string `db:"imgsrc"`
}

func ConvertProfileToTable(m models.Profile) ProfileTable {
	return ProfileTable{
		ID:          m.ID,
		Email:       m.Email,
		FullName:    m.FullName,
		PhoneNumber: m.PhoneNumber,
		ImgSrc:      m.ImgSrc,
	}
}

func ConvertTableToProfile(t ProfileTable) models.Profile {
	return models.Profile{
		ID:          t.ID,
		Email:       t.Email,
		FullName:    t.FullName,
		PhoneNumber: t.PhoneNumber,
		ImgSrc:      t.ImgSrc,
	}
}
