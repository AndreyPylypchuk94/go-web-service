package dao

import (
	"log"
	"pylypchuk.home/internal/model"
	"pylypchuk.home/pkg/store"
)

const (
	insertQuery        = "insert into app_user(email, first_name, last_name, password) values ($1, $2, $3, $4)"
	existsByEmailQuery = "select exists(select id from app_user where email = $1)"
	findByEmailQuery   = "select * from app_user where email = $1"
	findByIdQuery      = "select * from app_user where id = $1"
	deleteByIdQuery    = "delete from app_user where id = $1"
)

type user struct {
	dbClient *store.DbClient
}

func NewUserRepo(client *store.DbClient) *user {
	return &user{client}
}

func (u *user) FindAll() []model.User {
	var users []model.User
	err := u.dbClient.Db.Select(&users, "SELECT * FROM app_user")
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return users
}

func (u *user) Update(dto model.CreateUser) (int64, error) {
	var id int64
	err := u.dbClient.Db.QueryRow(insertQuery, dto.Email, dto.FirstName, dto.LastName, dto.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *user) Create(dto model.CreateUser) error {
	_, err := u.dbClient.Db.Query(insertQuery, dto.Email, dto.FirstName, dto.LastName, dto.Password)
	return err
}

func (u *user) ExistsByEmail(email string) bool {
	var exists bool
	_ = u.dbClient.Db.QueryRow(existsByEmailQuery, email).Scan(&exists)
	return exists
}

func (u *user) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.dbClient.Db.Get(&user, findByEmailQuery, email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) FindById(id int64) (*model.User, error) {
	var user model.User
	if err := u.dbClient.Db.Get(&user, findByIdQuery, id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) DeleteById(id int64) error {
	_, err := u.dbClient.Db.Exec(deleteByIdQuery, id)
	if err != nil {
		return err
	}
	return nil
}
