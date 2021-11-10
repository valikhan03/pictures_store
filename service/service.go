package service

import(
	"mime/multipart"
	"picturestore/repository"
)

type Upload interface{
	UploadFile(filename string, file multipart.File) error
}

type Service struct{
	Upload
}


func NewService(repository *repository.Repository) *Service{
	return &Service{
		Upload: NewUploadService(repository.Upload),
	}
}
