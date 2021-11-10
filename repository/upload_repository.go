package repository

import(
	"github.com/jmoiron/sqlx"
)

type UploadRepository struct{
	DB *sqlx.DB
}

func NewUploadRepository(db *sqlx.DB) *UploadRepository{
	return &UploadRepository{
		DB: db,
	}
}