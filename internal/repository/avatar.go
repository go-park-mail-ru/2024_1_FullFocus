package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
	"github.com/minio/minio-go/v7"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
)

type AvatarStorage struct {
	bucketName string
	client     *minio.Client
}

func NewAvatarStorage(c *minio.Client, cfg config.MinioConfig) *AvatarStorage {
	return &AvatarStorage{
		bucketName: cfg.AvatarBucket,
		client:     c,
	}
}

func (s *AvatarStorage) GetAvatar(ctx context.Context, fileName string) (models.Avatar, error) {
	start := time.Now()
	reader, err := s.client.GetObject(
		ctx,
		s.bucketName,
		fileName,
		minio.GetObjectOptions{})
	if err != nil {
		logger.Error(ctx, "get avatar error: "+err.Error())
		return models.Avatar{}, models.ErrNoAvatar
	}
	info, err := reader.Stat()
	if err != nil {
		logger.Error(ctx, "get file stat error: "+err.Error())
		return models.Avatar{}, models.ErrNoAvatar
	}
	avatar := models.Avatar{
		Payload:     reader,
		PayloadSize: info.Size,
	}
	logger.Info(ctx, fmt.Sprintf("%d bytes read in %s", info.Size, time.Since(start)))
	return avatar, nil
}

func (s *AvatarStorage) UploadAvatar(ctx context.Context, fileName string, img models.Avatar) error {
	start := time.Now()
	info, err := s.client.PutObject(
		ctx,
		s.bucketName,
		fileName,
		img.Payload,
		img.PayloadSize,
		minio.PutObjectOptions{})
	if err != nil {
		logger.Error(ctx, "upload avatar error: "+err.Error())
		return models.ErrCantUpload
	}
	logger.Info(ctx, fmt.Sprintf("%d bytes uploaded in %s: %s", info.Size, time.Since(start), fileName))
	return nil
}

func (s *AvatarStorage) DeleteAvatar(ctx context.Context, fileName string) error {
	start := time.Now()
	err := s.client.RemoveObject(
		ctx,
		s.bucketName,
		fileName,
		minio.RemoveObjectOptions{})
	if err != nil {
		logger.Error(ctx, "delete avatar error: "+err.Error())
		return models.ErrNoAvatar
	}
	logger.Info(ctx, fmt.Sprintf("file `%s` removed in %s", fileName, time.Since(start)))
	return err
}
