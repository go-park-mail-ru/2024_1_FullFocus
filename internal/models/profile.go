package models

type Profile struct {
	ID          uint   `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	ImgSrc      string `json:"imgsrc"`
}
