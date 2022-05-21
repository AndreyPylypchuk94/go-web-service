package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DbClient struct {
	Db *sqlx.DB
}

type DbConfig struct {
	User     string
	Password string
	DbName   string
}

const (
	dbConnectParams = "user=%s password=%s dbname=%s sslmode=disable"
	dbDriverName    = "postgres"
)

func NewDbClient(config DbConfig) *DbClient {
	dbConnectParamString := fmt.Sprintf(dbConnectParams, config.User, config.Password, config.DbName)
	db, err := sqlx.Connect(dbDriverName, dbConnectParamString)
	if err != nil {
		log.Fatalln(err)
	}
	return &DbClient{db}
}
