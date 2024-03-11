package models

type Profile struct {
	ProfileID   uint `json:"profieId"`
	User        User
	PhoneNumber string `json:"phoneNum"`
	Points      uint   `json:"points"`
	Img         string `json:"img-href"` // временное решение пока нет бд
}
