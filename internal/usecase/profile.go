package usecase

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type ProfileUsecase struct {
	userRepo repository.Users
}

func NewProfileUsecase(ur repository.Users) *ProfileUsecase {
	return &ProfileUsecase{
		userRepo: ur,
	}
}
