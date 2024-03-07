package repository

import (
	"fmt"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
)

type ChatListPostgres struct {
	db *sqlx.DB
}

func NewChatListPostgres(db *sqlx.DB) *ChatListPostgres {
	return &ChatListPostgres{db: db}
}

func (r *ChatListPostgres) Create(userId int, list chat.ChatList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", chatListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ChatListPostgres) GetAll(userId int) ([]chat.ChatList, error) {
	var lists []chat.ChatList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		chatListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *ChatListPostgres) GetById(userId, listId int) (chat.ChatList, error) {
	var list chat.ChatList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl 
						INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		chatListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}
