package repository

import "github.com/jmoiron/sqlx"


type Autorization interface {
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
	return &Repository{}
}
