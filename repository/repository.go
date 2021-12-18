package repository

import (
	"io"
	"picturestore/entity"

	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
)

type Auth interface{
	NewUser(userdata entity.User) error
	FindUser(userdata entity.SignInInput) (entity.User, error)
}

type Upload interface{
	UploadOne(user_id string, filename string, file io.Reader, size int64) error
}


type Repository struct {
	Auth
	Upload
}

func NewRepository(db *sqlx.DB, cl *minio.Client ) *Repository {
	return &Repository{
		Auth:NewAuthRepository(db),
		Upload: NewUploadRespository(cl),
	}
}
