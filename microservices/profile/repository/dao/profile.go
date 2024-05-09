package dao

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/models"
)

type ProfileTable struct {
	ID         uint   `db:"id"`
	FullName   string `db:"full_name"`
	AvatarName string `db:"imgsrc"`
}

func ConvertTableToProfile(t ProfileTable) models.Profile {
	return models.Profile{
		ID:         t.ID,
		FullName:   t.FullName,
		AvatarName: t.AvatarName,
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

type ProfileNameAvatar struct {
	FullName   string `db:"full_name"`
	AvatarName string `db:"imgsrc"`
}

func ConvertProfileNameAvatarToModel(t ProfileNameAvatar) models.ProfileNameAvatar {
	return models.ProfileNameAvatar{
		FullName:   t.FullName,
		AvatarName: t.AvatarName,
	}
}

func ConvertProfileNamesAvatarsToModels(tt []ProfileNameAvatar) []models.ProfileNameAvatar {
	data := make([]models.ProfileNameAvatar, 0)
	for _, t := range tt {
		data = append(data, ConvertProfileNameAvatarToModel(t))
	}
	return data
}
