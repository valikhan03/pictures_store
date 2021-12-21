package repository

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOConfigs struct {
	AccessKey string
	SecretKey string
	Endpoint  string
}

func NewMinIOStorage() *minio.Client {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file : %s", err.Error())
	}
	var configs = &MinIOConfigs{
		Endpoint:  os.Getenv("MINIO_ENDPOINT"),
		AccessKey: os.Getenv("MINIO_STORAGE_ACCESS_KEY"),
		SecretKey: os.Getenv("MINIO_STORAGE_SECRET_KEY"),
	}

	minioClient, err := minio.New(configs.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(configs.AccessKey, configs.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("MinIO client error - %x", err)
	}

	return minioClient
}
