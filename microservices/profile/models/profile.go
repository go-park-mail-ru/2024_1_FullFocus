package models

type Profile struct {
	ID         uint
	FullName   string
	AvatarName string
}

type ProfileMetaInfo struct {
	FullName   string
	AvatarName string
}

type ProfileUpdateInput struct {
	FullName string
}

type ProfileNameAvatar struct {
	FullName   string
	AvatarName string
}
