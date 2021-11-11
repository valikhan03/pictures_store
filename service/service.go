package service

import(
	"mime/multipart"
	"picturestore/repository"
	"picturestore/entity"
)

type Auth interface{
	SignUp(userdata entity.User) error
}

type Upload interface{
	UploadFile(filename string, file multipart.File) error
}

type Service struct{
	Upload
	Auth
}


func NewService(repository *repository.Repository) *Service{
	return &Service{
		Upload: NewUploadService(repository.Upload),
		Auth: NewAuthService(repository.Auth),
	}
}
