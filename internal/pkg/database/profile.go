package database

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type ProfileTable struct {
	ID          uint   `db:"id"`
	FullName    string `db:"full_name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	ImgSrc      string `db:"imgsrc"`
}

type ProfileUpdateTable struct {
	ID          uint   `db:"id"`
	FullName    string `db:"full_name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	ImgSrc      string `db:"imgsrc"`
}

func ConvertProfileToTable(m models.Profile) ProfileTable {
	return ProfileTable{
		ID:          m.ID,
		FullName:    m.FullName,
		Email:       m.Email,
		PhoneNumber: m.PhoneNumber,
		ImgSrc:      m.ImgSrc,
	}
}

func ConvertTableToProfile(t ProfileTable) models.Profile {
	return models.Profile{
		ID:          t.ID,
		FullName:    t.FullName,
		Email:       t.Email,
		PhoneNumber: t.PhoneNumber,
		ImgSrc:      t.ImgSrc,
	}
}
