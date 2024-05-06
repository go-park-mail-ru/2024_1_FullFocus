package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdatePasswordInput struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type SignupData struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
}

func ConvertSignupDataToModel(data SignupData) models.SignupData {
	return models.SignupData{
		Login:       data.Login,
		Password:    data.Password,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		FullName:    data.FullName,
	}
}
