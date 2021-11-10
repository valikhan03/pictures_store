package service

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"picturestore/repository"
)

type UploadService struct{
	repos repository.Upload
}

func NewUploadService(repos repository.Upload) *UploadService{
	return &UploadService{
		repos: repos,
	}
}

func (u *UploadService) UploadFile(filename string, file multipart.File) error {
	dst, err := os.Create("img_storage/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		log.Fatal(err)
	}

	return err
}