package usecase

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

const _avatarMaxSize = 1 << 21

type AvatarUsecase struct {
	avatarRepo  repository.Avatars
	profileRepo repository.Profiles
}

func NewAvatarUsecase(ar repository.Avatars, pr repository.Profiles) *AvatarUsecase {
	return &AvatarUsecase{
		avatarRepo:  ar,
		profileRepo: pr,
	}
}

func (u *AvatarUsecase) GetAvatar(ctx context.Context, profileID uint) (models.Avatar, error) {
	fileName, err := u.profileRepo.GetAvatarByProfileID(ctx, profileID)
	if err != nil {
		return models.Avatar{}, err
	}
	return u.avatarRepo.GetAvatar(ctx, fileName)
}

func (u *AvatarUsecase) UploadAvatar(ctx context.Context, profileID uint, img models.Avatar) error {
	if img.PayloadSize > _avatarMaxSize {
		return helper.NewValidationError("file too large", "Размер файла не должен превышать 2 MB")
	}
	file := bytes.NewBuffer(nil)
	_, err := io.Copy(file, img.Payload)
	if err != nil {
		return helper.NewValidationError("upload error", "Ошибка чтения файла")
	}
	fileType := http.DetectContentType(file.Bytes())
	switch fileType {
	case "image/jpeg", "image/jpg", "image/png":
		break
	default:
		return helper.NewValidationError("invalid file type", "Недопустимое расширение файла")
	}
	fileName := fmt.Sprintf("%d%d", profileID, time.Now().UnixNano())
	object := models.Avatar{
		Payload:     bytes.NewReader(file.Bytes()),
		PayloadSize: img.PayloadSize,
	}
	if err = u.avatarRepo.UploadAvatar(ctx, fileName, object); err != nil {
		return err
	}
	prevFileName, err := u.profileRepo.GetAvatarByProfileID(ctx, profileID)
	if err != nil {
		return err
	}
	if err = u.profileRepo.UpdateAvatarByProfileID(ctx, profileID, fileName); err != nil {
		return err
	}
	if prevFileName != "" {
		return u.avatarRepo.DeleteAvatar(ctx, prevFileName)
	}
	return nil
}

func (u *AvatarUsecase) DeleteAvatar(ctx context.Context, uID uint) error {
	fileName, err := u.profileRepo.GetAvatarByProfileID(ctx, uID)
	if err != nil {
		return err
	}
	if err = u.avatarRepo.DeleteAvatar(ctx, fileName); err != nil {
		return err
	}
	return u.profileRepo.DeleteAvatarByProfileID(ctx, uID)
}
