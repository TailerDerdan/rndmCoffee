package main

import (
	"os"

	chat "github.com/MerBasNik/rndmCoffee"
	handler "github.com/MerBasNik/rndmCoffee/pkg/handlers"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
	"github.com/MerBasNik/rndmCoffee/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := intiConfig(); err != nil {
		logrus.Fatalf("error initializing configs: $s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env value: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Password: os.Getenv("DB_PASSWORD"),
		Username: viper.GetString("db.Username"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialized db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(chat.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRouts()); err != nil {
		logrus.Fatalf("error while running http server: %s", err.Error())
	}
}

func intiConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
