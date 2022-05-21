package service

import (
	"pylypchuk.home/internal/dao"
	"pylypchuk.home/internal/model"
)

type UserWebService struct {
	userRepo dao.UserCrud
}

func NewUserWebService(userRepo dao.UserCrud) *UserWebService {
	return &UserWebService{userRepo: userRepo}
}

func (us *UserWebService) Get() []model.User {
	return us.userRepo.FindAll()
}

func (us *UserWebService) GetById(id int64) (*model.User, error) {
	return us.userRepo.FindById(id)
}

func (us *UserWebService) Update(dto model.CreateUser) (int64, error) {
	return us.userRepo.Update(dto)
}

func (us *UserWebService) DeleteById(id int64) error {
	return us.userRepo.DeleteById(id)
}
