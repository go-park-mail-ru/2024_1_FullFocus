package models

type Profile struct {
	ID         uint
	FullName   string
	AvatarName string
}

type ProfileMetaInfo struct {
	FullName        string
	CartItemsAmount uint
	AvatarName      string
}

type ProfileUpdateInput struct {
	FullName    string
	Email       string
	PhoneNumber string
}

type ProfileNameAvatar struct {
	FullName   string
	AvatarName string
}

type FullProfile struct {
	Email       string
	ProfileData Profile
}
