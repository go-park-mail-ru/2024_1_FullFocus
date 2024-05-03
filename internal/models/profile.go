package models

type Profile struct {
	ID          uint
	FullName    string
	Email       string
	PhoneNumber string
	AvatarName  string
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
