package service

import (
	"errors"
	"log"
	"picturestore/entity"
	"picturestore/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type AuthService struct {
	repos repository.Auth
}

func NewAuthService(repos repository.Auth) *AuthService {
	return &AuthService{
		repos: repos,
	}
}

const(
	tokenET = 5 * time.Minute
	key = "vfdnjkvbhzjlkjca;nbhcibguiv;hvj/vnal/"
)

type tokenClaims struct{
	
	jwt.StandardClaims
	UserID uuid.UUID `json:"user_id"`
}

func (a *AuthService) SignUp(userdata entity.User) error{
	err := a.repos.NewUser(userdata)
	if err != nil{
		log.Fatal(err)
	}
	return err
}




func (a *AuthService) GenerateToken(userdata entity.SignInInput) (string, error){
	user, err := a.repos.FindUser(userdata)
	if err != nil{
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenET).Unix(),
		},
		user.UserID,
	})
	
	tokenStr, err := token.SignedString(key)
	if err != nil{
		log.Fatal(err)
	}
	return tokenStr, err
}

func (a *AuthService) ParseToken(access_token string) (uuid.UUID, error){
	token, err := jwt.ParseWithClaims(access_token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("Invalid signing method")
		}
		return []byte(key), nil
	})

	if err != nil{
		log.Fatal(err)
	}

	cliams, ok := token.Claims.(*tokenClaims)
	if !ok {
		return uuid.Nil, errors.New("token claims error")
	}

	return cliams.UserID, nil
}