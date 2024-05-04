package dao

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/models"
)

type ProfileTable struct {
	ID          uint   `db:"id"`
	FullName    string `db:"full_name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	AvatarName  string `db:"imgsrc"`
}

func ConvertTableToProfile(t ProfileTable) models.Profile {
	return models.Profile{
		ID:          t.ID,
		FullName:    t.FullName,
		Email:       t.Email,
		PhoneNumber: t.PhoneNumber,
		AvatarName:  t.AvatarName,
	}
}

type ProfileMetaInfo struct {
	FullName   string `db:"full_name"`
	AvatarName string `db:"imgsrc"`
}

func ConvertProfileMetaInfo(info ProfileMetaInfo) models.ProfileMetaInfo {
	return models.ProfileMetaInfo{
		FullName:   info.FullName,
		AvatarName: info.AvatarName,
	}
}
