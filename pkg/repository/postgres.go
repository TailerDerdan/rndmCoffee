package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	userTable       = "users"
	chatListsTable  = "chat_lists"
	usersListsTable = "users_lists"
	chatItemsTable  = "chat_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Password string
	Username string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
