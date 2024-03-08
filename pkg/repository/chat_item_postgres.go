package repository

import (
	"fmt"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
)

type ChatItemPostgres struct {
	db *sqlx.DB
}

func NewChatItemPostgres(db *sqlx.DB) *ChatItemPostgres {
	return &ChatItemPostgres{db: db}
}

func (r *ChatItemPostgres) Create(listId int, item chat.ChatItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", chatItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	return itemId, tx.Commit()
}

func (r *ChatItemPostgres) GetAll(userId, listId int) ([]chat.ChatItem, error) {
	var items []chat.ChatItem
	query := fmt.Sprintf(`SELECT * FROM %s ti INNER JOIN %s li on li.item_id = ti.id 
							INNER JOIN %s ul on ul.list_id = li.list_id WHERE 
							li.list_id = $1 AND ul.user_id = $2`, chatItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}
