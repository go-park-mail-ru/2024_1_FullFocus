package dto

import model "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type ProfileData struct {
	ID          uint   `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	ImgSrc      string `json:"imgSrc"`
}

func ConvertProfileDataToProfile(profile model.Profile) ProfileData {
	return ProfileData{
		ID:          profile.ID,
		FullName:    profile.FullName,
		Email:       profile.Email,
		PhoneNumber: profile.PhoneNumber,
		ImgSrc:      profile.ImgSrc,
	}
}

func ConvertProfileToProfileData(profile ProfileData) model.Profile {
	return model.Profile{
		ID:          profile.ID,
		FullName:    profile.FullName,
		Email:       profile.Email,
		PhoneNumber: profile.PhoneNumber,
		ImgSrc:      profile.ImgSrc,
	}
}
