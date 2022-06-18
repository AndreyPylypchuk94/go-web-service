package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"pylypchuk.home/internal/dao"
	"pylypchuk.home/internal/model"
	"pylypchuk.home/internal/request"
	"pylypchuk.home/internal/service/userStorage"
	"pylypchuk.home/pkg/context"
	"strconv"
)

const tokenKey = "adadad"

type AuthWebService struct {
	userRepo dao.UserCrud
}

func NewAuthWebService() *AuthWebService {
	return &AuthWebService{
		userRepo: context.Get("userRepo").(dao.UserCrud),
	}
}

type Claims struct {
	jwt.StandardClaims
	Roles []string `json:"roles"`
}

func (as *AuthWebService) SignIn(request request.LoginRequest) (string, error) {
	user, err := as.userRepo.FindByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if user.Password != request.Password {
		return "", errors.New("incorrect email or password")
	}

	//simple token
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
	//	Subject: strconv.FormatInt(user.Id, 10),
	//})

	//custom token structure
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		jwt.StandardClaims{
			Subject: strconv.FormatInt(user.Id, 10),
		},
		[]string{"user"},
	})

	userStorage.Add(user.Id)

	return token.SignedString([]byte(tokenKey))
}

func (as *AuthWebService) SignUp(dto model.CreateUser) error {
	if as.userRepo.ExistsByEmail(dto.Email) {
		return errors.New("already exists")
	}
	return as.userRepo.Create(dto)
}

func (as *AuthWebService) ParseToken(tokenHeader string) (int64, []string, error) {
	token, err := jwt.ParseWithClaims(tokenHeader, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		return 0, nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return 0, nil, errors.New("claims not parsed")
	}

	i, err := strconv.ParseInt(claims.Subject, 0, 64)

	if err != nil {
		return 0, nil, err
	}

	return i, claims.Roles, nil
}
