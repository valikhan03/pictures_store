package service

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"log"
	"os"
	"picturestore/entity"
	"picturestore/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
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

	userdata.Password = hashPassword(userdata.Password)
	if err != nil{
		log.Fatal(err)
	}
	err = a.repos.NewUser(userdata)
	if err != nil {
		log.Fatal(err)
	}
	return id.String(), err
}

func (a *AuthService) GenerateToken(userdata entity.SignInInput) (string, error) {
	var err error
	userdata.Password = hashPassword(userdata.Password)

	fmt.Println("GENERATING TOKEN: ", userdata)

	if err != nil{
		log.Fatal(err)
	} 
	
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

	godotenv.Load("keys.env")
	key := os.Getenv("TOKEN_KEY")

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

		godotenv.Load("keys.env")
		key := os.Getenv("TOKEN_KEY")

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


func hashPassword(password string) string{
	godotenv.Load("keys.env")
	salt := os.Getenv("SALT")
	pwd := sha512.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(salt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	return password
}