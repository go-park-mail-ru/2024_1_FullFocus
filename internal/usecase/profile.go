package usecase

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type ProfileUsecase struct {
	userRepo    repository.Users
	profileRepo repository.Profiles
}

func NewProfileUsecase(ur repository.Users, pr repository.Profiles) *ProfileUsecase {
	return &ProfileUsecase{
		userRepo:    ur,
		profileRepo: pr,
	}
}
