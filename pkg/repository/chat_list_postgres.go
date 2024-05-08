package repository

import (
	"fmt"
	"strings"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ChatListPostgres struct {
	db *sqlx.DB
}

func NewChatListPostgres(db *sqlx.DB) *ChatListPostgres {
	return &ChatListPostgres{db: db}
}

func (r *ChatListPostgres) CreateList(requestCreateList chat.RequestCreateList) (int, string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, "", err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", chatListsTable)
	row := tx.QueryRow(createListQuery, requestCreateList.Title)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, "", err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, chatlists_id) VALUES ($1, $2)", usersChatListsTable)
	for i := 0; i < len(requestCreateList.UsersId); i++ {
		_, err = tx.Exec(createUsersListQuery, requestCreateList.UsersId[i], id)
		if err != nil {
			tx.Rollback()
			return 0, "", err
		}
	}

	return id, requestCreateList.Title, tx.Commit()
}

func (r *ChatListPostgres) GetAllLists(userId int) ([]chat.ChatList, error) {
	var lists []chat.ChatList

	query := fmt.Sprintf("SELECT tl.id, tl.title FROM %s tl INNER JOIN %s ul on tl.id = ul.chatlists_id WHERE ul.user_id = $1",
		chatListsTable, usersChatListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *ChatListPostgres) GetListById(userId, listId int) (chat.ChatList, error) {
	var list chat.ChatList

	query := fmt.Sprintf(`SELECT tl.id, tl.title FROM %s tl INNER JOIN %s ul on tl.id = ul.chatlists_id WHERE ul.user_id = $1 AND ul.chatlists_id = $2`,
		chatListsTable, usersChatListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *ChatListPostgres) DeleteList(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.chatlists_id AND ul.user_id=$1 AND ul.chatlists_id=$2",
		chatListsTable, usersChatListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *ChatListPostgres) UpdateList(userId, listId int, input chat.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.chatlists_id AND ul.chatlists_id=$%d AND ul.user_id=$%d",
		chatListsTable, setQuery, usersChatListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ChatListPostgres) FindByTime(userId int, input chat.FindUserInput) (int, error) {
	var id int

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Println("Ошибка в транзакцияхххх")
		return id, err
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s (user_id, start_day, end_day, start_time, end_time) VALUES ($1, $2, $3, $4, $5)", findUsersTable)
	_, err = tx.Exec(createListQuery, userId, input.StartDay, input.EndDay, input.StartTime, input.EndTime)
	if err != nil {
		tx.Rollback()
		fmt.Println("Ошибка в вставке в базу")
		return id, err
	}

	tx.Commit()

	query := fmt.Sprintf(`SELECT tl.user_id FROM %s tl WHERE
	((tl.start_day BETWEEN $1 AND $2) OR (tl.end_day BETWEEN $1 AND $2)) AND 
	((tl.start_time BETWEEN $3 AND $4) OR (tl.end_time BETWEEN $3 AND $4)) AND tl.user_id!=$5`,
		findUsersTable)
	err = r.db.Get(&id, query, input.StartDay, input.EndDay, input.StartTime, input.EndTime, userId)

	return id, err
}

func (r *ChatListPostgres) FindByHobby(userId1, userId2 int) ([]chat.UserHobby, error) {
	var lists []chat.UserHobby

	query := fmt.Sprintf(`SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id
	WHERE ul.user_id=$1 INTERSECT SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id
	WHERE ul.user_id=$2`,
		userHobbyTable, usersHobbyListsTable, userHobbyTable, usersHobbyListsTable)
	err := r.db.Select(&lists, query, userId1, userId2)

	return lists, err
}
