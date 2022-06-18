package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"pylypchuk.home/internal/api"
	dao "pylypchuk.home/internal/dao/impl"
	"pylypchuk.home/internal/service"
	"pylypchuk.home/pkg"
	"pylypchuk.home/pkg/context"
	"pylypchuk.home/pkg/store"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	dbConfig := store.DbConfig{
		DbName:   viper.GetString("db.dbname"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	initContext(dbConfig)

	server := pkg.NewServer()
	if err := server.Run(viper.GetString("port")); err != nil {
		log.Fatal(err.Error())
	}
}

func initContext(dbConfig store.DbConfig) {
	context.Add("dbClient", store.NewDbClient(dbConfig))
	context.Add("userRepo", dao.NewUserRepo())
	context.Add("authWebService", service.NewAuthWebService())
	context.Add("userWebService", service.NewUserWebService())
	context.Add("handler", api.NewHandler())
}

func initConfig() error {
	viper.AddConfigPath("_config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
