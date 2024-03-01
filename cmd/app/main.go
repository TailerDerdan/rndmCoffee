package main

import (
	"log"
	"github.com/MerBasNik/endmCoffee"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while runnung http server: %s", err.Error())
	}
}
