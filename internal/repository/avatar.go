package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
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

func (s *AvatarStorage) UploadAvatar(ctx context.Context, img models.Image) error {
	l := helper.GetLoggerFromContext(ctx)
	start := time.Now()

	info, err := s.client.PutObject(
		ctx,
		_avatarBucket,
		img.Name,
		img.Payload,
		img.PayloadSize,
		minio.PutObjectOptions{})

	l.Info(fmt.Sprintf("%d bytes uploaded in %s", info.Size, time.Since(start)))
	return err
}

func (s *AvatarStorage) DeleteAvatar(ctx context.Context, imgName string) error {
	l := helper.GetLoggerFromContext(ctx)
	start := time.Now()

	err := s.client.RemoveObject(
		ctx,
		_avatarBucket,
		imgName,
		minio.RemoveObjectOptions{})

	l.Info(fmt.Sprintf("file removed in %s", time.Since(start)))
	return err
}
