package repository

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user chat.User) (int, error)
	GetUser(username, password string) (chat.User, error)
}

type ChatList interface {
	Create(userId int, list chat.ChatList) (int, error)
	GetAll(userId int) ([]chat.ChatList, error)
	GetById(userId, listId int) (chat.ChatList, error)
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
		ChatList:     NewChatListPostgres(db),
	}
}
