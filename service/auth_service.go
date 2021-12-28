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

const (
	tokenET = 60 * time.Minute //token expiration time
	key     = "vfdnjkvbhzjlkj"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func (a *AuthService) SignUp(userdata entity.User) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return "", err
	}
	userdata.UserID = id.String()
	err = a.repos.NewUser(userdata)
	if err != nil {
		log.Fatal(err)
	}
	return id.String(), err
}

func (a *AuthService) GenerateToken(userdata entity.SignInInput) (string, error) {
	user, err := a.repos.FindUser(userdata)
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenET).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.UserID,
	})

	tokenStr, err := token.SignedString([]byte(key))
	if err != nil {
		log.Fatal("KEY ERROR", err)
	}
	return tokenStr, err
}

func (a *AuthService) ParseToken(access_token string) (string, error) {
	token, err := jwt.ParseWithClaims(access_token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(key), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	cliams, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims error")
	}

	return cliams.UserID, nil
}
