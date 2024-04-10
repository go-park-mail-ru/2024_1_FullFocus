package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
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

func (s *AvatarStorage) GetAvatar(ctx context.Context) models.Image {
	return models.Image{} // TODO
}

func (s *AvatarStorage) UploadAvatar(ctx context.Context, img models.Image) error {
	start := time.Now()
	info, err := s.client.PutObject(
		ctx,
		_avatarBucket,
		img.Name,
		img.Payload,
		img.PayloadSize,
		minio.PutObjectOptions{})
	logger.Info(ctx, fmt.Sprintf("%d bytes uploaded in %s", info.Size, time.Since(start)))
	return err
}

func (s *AvatarStorage) DeleteAvatar(ctx context.Context, imgName string) error {
	start := time.Now()
	err := s.client.RemoveObject(
		ctx,
		_avatarBucket,
		imgName,
		minio.RemoveObjectOptions{})
	logger.Info(ctx, fmt.Sprintf("file removed in %s", time.Since(start)))
	return err
}
