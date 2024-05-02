package dto

import model "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type Profile struct {
	ID          uint   `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	AvatarName  string `json:"avatarName"`
}

func ConvertProfileDataToProfile(profile model.Profile) Profile {
	return Profile{
		ID:          profile.ID,
		FullName:    profile.FullName,
		Email:       profile.Email,
		PhoneNumber: profile.PhoneNumber,
		AvatarName:  profile.AvatarName,
	}
}

type ProfileUpdateInput struct {
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func ConvertProfileToProfileData(profile ProfileUpdateInput) model.ProfileUpdateInput {
	return model.ProfileUpdateInput{
		FullName:    profile.FullName,
		Email:       profile.Email,
		PhoneNumber: profile.PhoneNumber,
	}
}
