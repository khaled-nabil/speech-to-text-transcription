package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type (
	MinIOConfig struct {
		Endpoint        string
		AccessKeyID     string
		SecretAccessKey string
		UseSSL          bool
		BucketName      string
	}

	StorageConfig struct {
		MaxFileSize int
		AllowedExt  []string
	}

	FasterWhisper struct {
		Endpoint string
		Port     string
		Model    string
	}

	Config struct {
		Env           string
		MinIO         MinIOConfig
		Storage       StorageConfig
		FasterWhisper FasterWhisper
	}
)

func New() (*Config, error) {
	c := &Config{
		Env: getEnvDefault("APP_ENV", "development"),
		MinIO: MinIOConfig{
			Endpoint:        getEnvFatal("MINIO_ENDPOINT"),
			AccessKeyID:     getEnvFatal("MINIO_ROOT_USER"),
			SecretAccessKey: getEnvFatal("MINIO_ROOT_PASSWORD"),
			UseSSL:          getBoolEnvFatal("MINIO_USE_SSL"),
			BucketName:      getEnvFatal("MINIO_BUCKET_NAME"),
		},
		Storage: StorageConfig{
			MaxFileSize: getEnvIntegerFatal("MAX_MB_FILE_SIZE"),
			AllowedExt:  getEnvListStringsFatal("ALLOWED_FILE_TYPES"),
		},
		FasterWhisper: FasterWhisper{
			Endpoint: getEnvFatal("FASTER_WHISPER_ENDPOINT"),
			Port:     getEnvFatal("FASTER_WHISPER_PORT"),
			Model:    getEnvFatal("FASTER_WHIPSER_MODEL"),
		},
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}

func getBoolEnvFatal(key string) bool {
	val := os.Getenv(key)
	if val == "" {
		panic("missing required env var: " + key)
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		panic("invalid bool value for env var: " + key)
	}
	return b
}

func getEnvFatal(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("missing required env var: " + key)
	}
	return val
}

func getEnvDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func getEnvIntegerFatal(key string) int {
	val := os.Getenv(key)
	if val == "" {
		return 0
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		panic("invalid integer value for env var: " + key)
	}
	return i
}

func getEnvListStringsFatal(key string) []string {
	val := os.Getenv(key)
	if val == "" {
		return nil
	}
	return strings.Split(val, ",")
}

func (c *Config) Validate() error {
	if c.MinIO.Endpoint == "" {
		return errors.New("minio endpoint required")
	}
	return nil
}
