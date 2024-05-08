package models

type Profile struct {
	ID          uint
	FullName    string
	Email       string
	PhoneNumber string
	AvatarName  string
}

type ProfileMetaInfo struct {
	FullName   string
	AvatarName string
}

type ProfileUpdateInput struct {
	FullName    string
	Email       string
	PhoneNumber string
}

type ProfileNameAvatar struct {
	ID         uint
	FullName   string
	AvatarName string
}
