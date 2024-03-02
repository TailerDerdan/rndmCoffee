package main

import (
	"log"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/handler"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
	"github.com/MerBasNik/rndmCoffee/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(chat.Server)
	if err := srv.Run("8000", handler.InitRouts()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
