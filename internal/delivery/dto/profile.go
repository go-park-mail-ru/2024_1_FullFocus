package dto

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type Profile struct {
	ID          uint   `json:"id"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNum"`
	Gender      uint   `json:"gender"`
	AvatarName  string `json:"avatarName"`
}

func ConvertProfileDataToProfile(profile models.Profile) Profile {
	return Profile{
		ID:          profile.ID,
		FullName:    profile.FullName,
		Address:     profile.Address,
		PhoneNumber: profile.PhoneNum,
		Gender:      profile.Gender,
		AvatarName:  profile.AvatarName,
	}
}

type ProfileUpdateInput struct {
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNum"`
	Gender      uint   `json:"gender"`
}

func ConvertProfileToProfileData(profile ProfileUpdateInput) models.ProfileUpdateInput {
	return models.ProfileUpdateInput{
		FullName: profile.FullName,
		Address:  profile.Address,
		PhoneNum: profile.PhoneNumber,
		Gender:   profile.Gender,
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
	Email       string `json:"email"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNum"`
	Gender      uint   `json:"gender"`
	AvatarName  string `json:"avatarName"`
}

func ConvertFullProfileDataToDto(m models.FullProfile) FullProfile {
	return FullProfile{
		Email:       m.Email,
		ID:          m.ProfileData.ID,
		FullName:    m.ProfileData.FullName,
		Address:     m.ProfileData.Address,
		PhoneNumber: m.ProfileData.PhoneNum,
		Gender:      m.ProfileData.Gender,
		AvatarName:  m.ProfileData.AvatarName,
	}
}
