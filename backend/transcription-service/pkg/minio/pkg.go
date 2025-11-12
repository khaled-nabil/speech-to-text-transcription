package minio

import (
	"bytes"
	"context"
	"io"
	"transcription-service/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	client *minio.Client
	conf   *config.MinIOConfig
}

func New(conf *config.MinIOConfig) (*Storage, error) {
	minioClient, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKeyID, conf.SecretAccessKey, ""),
		Secure: conf.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	return &Storage{client: minioClient, conf: conf}, nil
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
	defer object.Close()

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
