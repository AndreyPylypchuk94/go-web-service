package model

type User struct {
	Id        int64  `json:"id" required:"true"`
	Email     string `json:"email" required:"true"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Password  string `json:"password" required:"true"`
}

type CreateUser struct {
	Email     string `json:"email" required:"true"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Password  string `json:"password" required:"true"`
}
