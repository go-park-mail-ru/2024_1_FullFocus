package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type AvatarUsecase struct {
	avatarRepo repository.Avatars
	userRepo   repository.Users // TODO: repository.Profiles
}

func NewAvatarUsecase(ar repository.Avatars, ur repository.Users) *AvatarUsecase {
	return &AvatarUsecase{
		avatarRepo: ar,
		userRepo:   ur,
	}
}

func (u *AvatarUsecase) UploadAvatar(ctx context.Context, img dto.Image) error {
	uID, _ := helper.GetUserIDFromContext(ctx)
	imgName := fmt.Sprintf("%d_%d", uID, time.Now().UnixNano())

	object := models.Image{
		Name:        imgName,
		Payload:     img.Payload,
		PayloadSize: img.PayloadSize,
	}
	// TODO: удалить прежнюю аву, если есть
	if err := u.avatarRepo.UploadAvatar(ctx, object); err != nil {
		return err
	}
	// TODO: обновить ссылку на аву пользователя с id = uid
	return nil
}

func (u *AvatarUsecase) DeleteAvatar(ctx context.Context) error {
	uID, _ := helper.GetUserIDFromContext(ctx)
	// TODO: получить имя авы по id пользователя  |
	//                                            | вроде бы можно одним запросом с `RETURNING`
	// TODO: удалить имя авы у этого пользователя |
	// TODO: если имя авы пустое, то вернуть models.ErrNoAvatar

	avatarName := fmt.Sprintf("avatar_%d", uID)
	return u.avatarRepo.DeleteAvatar(ctx, avatarName)
}
