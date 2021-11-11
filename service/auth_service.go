package service

import (
	"log"
	"picturestore/entity"
	"picturestore/repository"
)

type AuthService struct {
	repos repository.Auth
}

func NewAuthService(repos repository.Auth) *AuthService {
	return &AuthService{
		repos: repos,
	}
}


func (a *AuthService) SignUp(userdata entity.User) error{
	err := a.repos.NewUser(userdata)
	if err != nil{
		log.Fatal(err)
	}
	return err
}




func GenerateToken(){}

func ParseToken(){}