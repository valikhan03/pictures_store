package repository

import (
	"log"

	"github.com/minio/minio-go/v7"
)

type MinIOConfigs struct {
	Endpoint string
	UseSSL   bool
}

func NewMinIOStorage() *minio.Client{
	

	var configs MinIOConfigs

	minioClient, err := minio.New(configs.Endpoint, &minio.Options{Secure: configs.UseSSL})
	if err != nil{
		log.Fatal(err)
	}

	return minioClient
}
