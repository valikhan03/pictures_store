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

func (u *UploadService) UploadFile(userid string, filename string, file multipart.File, size int64) error {
	dst, err := os.Create("img_storage/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		log.Fatal(err)
	}

	err = u.repos.UploadOne(userid, filename, file, size)

	return err
}