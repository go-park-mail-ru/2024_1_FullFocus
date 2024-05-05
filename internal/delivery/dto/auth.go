package dto

type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdatePasswordInput struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
