package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"pylypchuk.home/internal/dao"
	"pylypchuk.home/internal/model"
	"pylypchuk.home/internal/request"
	"strconv"
)

const tokenKey = "adadad"

type AuthWebService struct {
	userRepo dao.UserCrud
}

func NewAuthWebService(userRepo dao.UserCrud) *AuthWebService {
	return &AuthWebService{userRepo: userRepo}
}

type Claims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

func (as *AuthWebService) SignIn(request request.LoginRequest) (string, error) {
	user, err := as.userRepo.FindByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if user.Password != request.Password {
		return "", errors.New("incorrect email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Subject: strconv.FormatInt(user.Id, 10),
	})

	return token.SignedString([]byte(tokenKey))
}

func (as *AuthWebService) SignUp(dto model.CreateUser) error {
	if as.userRepo.ExistsByEmail(dto.Email) {
		return errors.New("already exists")
	}
	return as.userRepo.Create(dto)
}

func (as *AuthWebService) ParseToken(tokenHeader string) (int64, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenHeader, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(claims["sub"].(string), 0, 64)

	if err != nil {
		return 0, err
	}

	return i, nil
}
