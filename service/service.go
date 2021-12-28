package service

import(
	"mime/multipart"
	"picturestore/repository"
	"picturestore/entity"

)

type Auth interface{
	SignUp(userdata entity.User) (string, error)
	GenerateToken(userdata entity.SignInInput) (string, error)
	ParseToken(access_token string) (string, error)
}

type Storage interface{
	UploadFile(userid string, filename string, file multipart.File, size int64) error
	NewUserBucket(user_id string) error 
	GetFile(user_id string, filename string) ([]byte, error)
	GetAllFilesList(user_id string)
}

type Service struct{
	Storage
	Auth
}


func NewService(repository *repository.Repository) *Service{
	return &Service{
		Storage: NewStorageService(repository.Storage),
		Auth: NewAuthService(repository.Auth),
	}
}
