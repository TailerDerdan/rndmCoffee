package main

import (
	"log"

	_ "github.com/lib/pq"
	chat "github.com/MerBasNik/rndmCoffee"
	handler "github.com/MerBasNik/rndmCoffee/pkg/handlers"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
	"github.com/MerBasNik/rndmCoffee/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := intiConfig(); err != nil {
		log.Fatalf("error initializing configs: $s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Password: "qwerty",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatalf("failed to initialized db: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(chat.Server)
	if err := srv.Run(viper.GetString("8000"), handler.InitRouts()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}

func intiConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
