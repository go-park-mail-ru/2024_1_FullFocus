package dto

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type Profile struct {
	ID          uint   `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	AvatarName  string `json:"avatarName"`
}

func ConvertProfileDataToProfile(profile models.Profile) Profile {
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

func ConvertProfileToProfileData(profile ProfileUpdateInput) models.ProfileUpdateInput {
	return models.ProfileUpdateInput{
		FullName:    profile.FullName,
		Email:       profile.Email,
		PhoneNumber: profile.PhoneNumber,
	}
}

type ProfileMetaInfo struct {
	FullName        string `json:"fullName"`
	CartItemsAmount uint   `json:"cartItemsAmount"`
	AvatarName      string `json:"avatarName"`
}

func ConvertProfileToMetaInfo(profile models.ProfileMetaInfo) ProfileMetaInfo {
	return ProfileMetaInfo{
		FullName:        profile.FullName,
		CartItemsAmount: profile.CartItemsAmount,
		AvatarName:      profile.AvatarName,
	}
}

type FullProfile struct {
	ID          uint   `json:"id"`
	Login       string `json:"login"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	AvatarName  string `json:"avatarName"`
}

func ConvertFullProfileDataToDto(m models.FullProfile) FullProfile {
	return FullProfile{
		ID:          m.ProfileData.ID,
		FullName:    m.ProfileData.FullName,
		Email:       m.ProfileData.Email,
		PhoneNumber: m.ProfileData.PhoneNumber,
		AvatarName:  m.ProfileData.AvatarName,
		Login:       m.Login,
	}
}
