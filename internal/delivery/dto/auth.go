package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdatePasswordInput struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type SignupData struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

func ConvertSignupDataToModel(data SignupData) models.SignupData {
	return models.SignupData{
		Password: data.Password,
		Email:    data.Email,
		FullName: data.FullName,
	}
}
