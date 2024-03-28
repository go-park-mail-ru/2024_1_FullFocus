package repository

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/minio/minio-go/v7"
	"time"
)

const (
	_avatarBucket = "avatar"
)

type AvatarStorage struct {
	client *minio.Client
}

func NewAvatarStorage(c *minio.Client) *AvatarStorage {
	return &AvatarStorage{
		client: c,
	}
}

func (s *AvatarStorage) UploadAvatar(ctx context.Context, img models.ImageUnit) error {
	log := helper.GetLoggerFromContext(ctx)
	start := time.Now()

	info, err := s.client.PutObject(
		ctx,
		_avatarBucket,
		img.Name,
		img.Payload,
		img.PayloadSize,
		minio.PutObjectOptions{})

	log.Info(fmt.Sprintf("%d bytes uploaded in %s", info.Size, time.Since(start)))
	return err
}

func (s *AvatarStorage) DeleteAvatar(ctx context.Context, imageName string) error {
	log := helper.GetLoggerFromContext(ctx)
	start := time.Now()

	err := s.client.RemoveObject(
		ctx,
		_avatarBucket,
		imageName,
		minio.RemoveObjectOptions{})

	log.Info(fmt.Sprintf("file removed in %s", time.Since(start)))
	return err
}
