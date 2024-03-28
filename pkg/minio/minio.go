package minio

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
)

func NewClient(cfg config.MinioConfig) (*minio.Client, error) {
	return minio.New(fmt.Sprintf("127.0.0.1%s", cfg.Port), &minio.Options{
		Creds: credentials.NewStaticV4(cfg.MinioUser, cfg.MinioPassword, ""),
	})
}
