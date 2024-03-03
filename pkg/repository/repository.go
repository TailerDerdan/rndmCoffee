package repository

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user chat.User) (int, error)
}

type ChatList interface {
}

type ChatItem interface {
}

type Repository struct {
	Autorization
	ChatItem
	ChatList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}
