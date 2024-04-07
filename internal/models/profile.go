package models

type Profile struct {
	User  User  `json:"user"`
	Image Image `json:"image"`
}
