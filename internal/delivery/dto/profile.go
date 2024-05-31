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

type ProfileMetaInfo struct {
	FullName            string `json:"fullName"`
	CartItemsAmount     uint   `json:"cartItemsAmount"`
	AvatarName          string `json:"avatarName"`
	UnreadNotifications uint   `json:"unreadNotifications"`
	PromocodesAvailable uint   `json:"promocodesAvailable"`
}

func ConvertProfileToMetaInfo(profile model.ProfileMetaInfo) ProfileMetaInfo {
	return ProfileMetaInfo{
		FullName:            profile.FullName,
		CartItemsAmount:     profile.CartItemsAmount,
		AvatarName:          profile.AvatarName,
		UnreadNotifications: profile.UnreadNotifications,
		PromocodesAvailable: profile.PromocodesAvailable,
	}
}

type FullProfile struct {
	Login       string `json:"login"`
	ProfileData Profile
}

func ConvertFullProfileToDto(m model.FullProfile) FullProfile {
	return FullProfile{
		Login:       m.Login,
		ProfileData: ConvertProfileDataToProfile(m.ProfileData),
	}
}
