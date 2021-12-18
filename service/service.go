package service

import(
	"mime/multipart"
	"picturestore/repository"
	"picturestore/entity"

	"github.com/google/uuid"
)

type Auth interface{
	SignUp(userdata entity.User) error
	GenerateToken(userdata entity.SignInInput) (string, error)
	ParseToken(access_token string) (uuid.UUID, error)
}

type Upload interface{
	UploadFile(userid string, filename string, file multipart.File, size int64) error
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
