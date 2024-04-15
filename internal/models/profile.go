package models

type Profile struct {
	ID          uint
	FullName    string
	Email       string
	PhoneNumber string
}

type ProfileUpdateInput struct {
	FullName    string
	Email       string
	PhoneNumber string
}
