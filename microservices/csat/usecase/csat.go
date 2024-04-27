package usecase

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/models"
)

type CSATs interface {
	GetAllPolls()
	AddPollRate()
}

type CSATUsecase struct {
	CSATRepo CSATs
}

func GetAllPolls() ([]models.Poll, error) {
	return
}
