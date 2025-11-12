package minio

import (
	"bytes"
	"context"
	"io"
	"log"
	"transcription-service/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	client *minio.Client
	conf   *config.MinIOConfig
}

func New(c *config.Config) (*Storage, error) {
	minioConfig := c.MinIO
	minioClient, err := minio.New(minioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioConfig.AccessKeyID, minioConfig.SecretAccessKey, ""),
		Secure: minioConfig.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	// Check if the bucket exists, create if not
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, minioConfig.BucketName)
	if err != nil {
		log.Printf("Error checking if bucket %s exists: %s\n", minioConfig.BucketName, err)
		return nil, err
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, minioConfig.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return &Storage{client: minioClient, conf: &minioConfig}, nil
}

func (s *Storage) StoreFile(path string, data []byte) error {
	ctx := context.Background()

	_, err := s.client.PutObject(ctx, s.conf.BucketName, path,
		bytes.NewReader(data), int64(len(data)),
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		})

	return err
}

func (s *Storage) GetFile(path string) ([]byte, error) {
	ctx := context.Background()

	object, err := s.client.GetObject(ctx, s.conf.BucketName, path, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer func(object *minio.Object) {
		err = object.Close()
		if err != nil {
			log.Printf("Error closing object: %v", err)
		}
	}(object)

	data, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Storage) DeleteFile(path string) error {
	ctx := context.Background()

	err := s.client.RemoveObject(ctx, s.conf.BucketName, path, minio.RemoveObjectOptions{})

	return err
}
