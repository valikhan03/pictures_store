package repository

import (
	"log"
	"picturestore/entity"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (a *AuthRepository) NewUser(userdata entity.User) error {
	_, err := a.DB.Exec("insert into users (id, email, username, password) values ($1, $2, $3, $4)",
		userdata.UserID, userdata.Email, userdata.Username, userdata.Password)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (a *AuthRepository) FindUser(userdata entity.SignInInput) (entity.User, error) {
	var user entity.User
	query := "select * from users where email=$1 and password=$2"

	err := a.DB.Get(&user, query, userdata.Email, userdata.Password)
	if err != nil {
		log.Println(err)
	}

	return user, err
}
