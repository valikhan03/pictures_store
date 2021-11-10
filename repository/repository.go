package repository

import (
	"github.com/jmoiron/sqlx"
)

type Upload interface{

}

type Repository struct {
	Upload
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Upload: NewUploadRepository(db),
	}
}
