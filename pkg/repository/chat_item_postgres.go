package repository

import (
	"fmt"
	"strconv"

	// "strings"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
)

type ChatItemPostgres struct {
	db *sqlx.DB
}

func NewChatItemPostgres(db *sqlx.DB) *ChatItemPostgres {
	return &ChatItemPostgres{db: db}
}

func (r *ChatItemPostgres) GetUsers(userId, chatId int) ([]int, error) {
	var usersID []int
	query := fmt.Sprintf("SELECT tl.user_id FROM %s tl WHERE tl.chatlists_id = $1", usersChatListsTable)
	if err := r.db.Select(&usersID, query, chatId); err != nil {
		return usersID, err
	}

	return usersID, nil
}

func (r *ChatItemPostgres) Create(chatId int, item chat.ChatItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Println("ТУууут????")
		return 0, err
	}

	fmt.Println(item.User_id)
	clientIdNum, err := strconv.Atoi(item.User_id)
	if err != nil {
		fmt.Println("Я ттттуттт????")
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (username, description, chatlist_id, user_id) values ($1, $2, $3, $4) RETURNING id", chatItemsTable)

	row := tx.QueryRow(createItemQuery, item.Username, item.Description, item.Chatlist_id, clientIdNum)
	err = row.Scan(&itemId)
	if err != nil {
		fmt.Println("может ТУТ????")
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *ChatItemPostgres) GetAll(userId, listId int) ([]chat.ChatItem, error) {
	var items []chat.ChatItem
	query := fmt.Sprintf(`SELECT ti.id, ti.username, ti.description, ti.chatlist_id, ti.user_id FROM %s ti 
	INNER JOIN %s ul on ul.chatlists_id = ti.chatlist_id WHERE ti.chatlist_id = $1 AND ul.user_id = $2`,
		chatItemsTable, usersChatListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		fmt.Println("Или тут????")
		return nil, err
	}

	return items, nil
}

// func (r *ChatItemPostgres) GetById(userId, itemId int) (chat.ChatItem, error) {
// 	var item chat.ChatItem
// 	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description FROM %s ti INNER JOIN %s li on li.chatitems_id = ti.id
// 	INNER JOIN %s ul on ul.chatlists_id = li.chatlists_id WHERE ti.id = $1 AND ul.user_id = $2`,
// 		chatItemsTable, itemsListsTable, usersChatListsTable)
// 	if err := r.db.Get(&item, query, itemId, userId); err != nil {
// 		return item, err
// 	}

// 	return item, nil
// }

// func (r *ChatItemPostgres) Delete(userId, itemId int) error {
// 	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
// 	WHERE ti.id = li.chatitems_id AND li.chatlists_id = ul.chatlists_id AND ul.user_id = $1 AND ti.id = $2`,
// 		chatItemsTable, itemsListsTable, usersChatListsTable)
// 	_, err := r.db.Exec(query, userId, itemId)
// 	return err
// }

// func (r *ChatItemPostgres) Update(userId, itemId int, input chat.UpdateItemInput) error {
// 	setValues := make([]string, 0)
// 	args := make([]interface{}, 0)
// 	argId := 1

// 	if input.Title != nil {
// 		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
// 		args = append(args, *input.Title)
// 		argId++
// 	}

// 	if input.Description != nil {
// 		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
// 		args = append(args, *input.Description)
// 		argId++
// 	}

// 	setQuery := strings.Join(setValues, ", ")

// 	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
// 						WHERE ti.id = li.chatitems_id AND li.chatlists_id = ul.chatlists_id AND ul.user_id = $%d AND ti.id = $%d`,
// 		chatItemsTable, setQuery, itemsListsTable, usersChatListsTable, argId, argId+1)
// 	args = append(args, userId, itemId)

// 	_, err := r.db.Exec(query, args...)
// 	return err
// }
