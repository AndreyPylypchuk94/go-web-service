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

	dbClient := store.NewDbClient(dbConfig)
	userRepo := dao.NewUserRepo(dbClient)
	authService := service.NewAuthWebService(userRepo)
	userService := service.NewUserWebService(userRepo)
	handler := api.NewHandler(userService, authService)

	server := new(pkg.Server)
	if err := server.Run(viper.GetString("port"), handler.InitRouts()); err != nil {
		log.Fatal(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("_config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
