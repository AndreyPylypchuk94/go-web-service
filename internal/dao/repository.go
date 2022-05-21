package dao

import "pylypchuk.home/internal/model"

type UserCrud interface {
	FindAll() []model.User
	Update(dto model.CreateUser) (int64, error)
	Create(dto model.CreateUser) error
	ExistsByEmail(email string) bool
	FindByEmail(email string) (*model.User, error)
	FindById(id int64) (*model.User, error)
	DeleteById(id int64) error
}
