package repository

import (
	"io"
	"picturestore/entity"

	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
)

type Auth interface {
	NewUser(userdata entity.User) error
	FindUser(userdata entity.SignInInput) (entity.User, error)
}

type Storage interface {
	GetFile(user_id, filename string) ([]byte, error)
	UploadOne(user_id string, filename string, file io.Reader, size int64) error
	MakeBucket(user_id string) error
	GetAllFilesList(user_id string) []string
}

type Repository struct {
	Auth
	Storage
}

func NewRepository(db *sqlx.DB, cl *minio.Client) *Repository {
	return &Repository{
		Auth:    NewAuthRepository(db),
		Storage: NewStorageRepository(cl),
	}
}
