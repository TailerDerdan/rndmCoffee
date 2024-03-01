package main

import (
	"log"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/handler"
)

func main() {
	handler := new(handler.Handler)

	srv := new(chat.Server)
	if err := srv.Run("8000", handler.InitRouts()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
