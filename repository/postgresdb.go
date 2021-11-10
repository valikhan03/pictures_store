package repository

import (
	"fmt"

	"picturestore/configs"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB() (*sqlx.DB, error) {
	var confs configs.PostgresDBConfigs
	confs = configs.GetPostgresDBConfigs()
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		confs.Host, confs.Port, confs.DBName, confs.Username, confs.Password, confs.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
