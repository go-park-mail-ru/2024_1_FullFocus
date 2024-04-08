package dto

type ProfileData struct {
	ID          uint   `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	ImgSrc      string `json:"imgSrc"`
}
