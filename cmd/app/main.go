package main

import (
	"log"

	todo "github.com/MerBasNik/rndmCoffee"
)

func main()  {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}