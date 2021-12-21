package service

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"picturestore/repository"
)

type StorageService struct {
	repos repository.Storage
}

func NewStorageService(repos repository.Storage) *StorageService {
	return &StorageService{
		repos: repos,
	}
}

func (u *StorageService) NewUserBucket(user_id string) error {
	err := u.repos.MakeBucket(user_id)
	return err
}

func (u *StorageService) UploadFile(userid string, filename string, file multipart.File, size int64) error {
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

func (s *StorageService) GetFile(user_id string, filename string) ([]byte, error){
	filedata, err := s.repos.GetFile(user_id, filename)
	return filedata, err
}