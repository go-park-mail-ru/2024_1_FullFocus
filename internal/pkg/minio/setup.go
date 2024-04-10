package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func InitBucket(ctx context.Context, client *minio.Client, bucketName string) error {
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		return client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}
	return nil
}
