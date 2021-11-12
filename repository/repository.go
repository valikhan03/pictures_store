package repository

import (
	"picturestore/entity"
	"github.com/jmoiron/sqlx"
)

type Auth interface{
	NewUser(userdata entity.User) error
	FindUser(userdata entity.SignInInput) (entity.User, error)
}

type Upload interface{

}

type Repository struct {
	Auth
	Upload
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
		Upload: NewUploadRepository(db),
	}
}
