package models

type Profile struct {
	ID          uint
	FullName    string
	Address     string
	PhoneNumber string
	Gender      uint
	AvatarName  string
}

type ProfileMetaInfo struct {
	FullName   string
	AvatarName string
}

type ProfileUpdateInput struct {
	FullName    string
	Address     string
	PhoneNumber string
	Gender      uint
}

type ProfileNameAvatar struct {
	FullName   string
	AvatarName string
}
