package chat

import (
	"errors"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Chats      map[string]*ChatList
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *ChatItem
}

type ChatList struct {
	Id      string             `json:"id" db:"id"`
	UsersId map[string]*Client `json:"clients"`
	Title   string             `json:"title" db:"title" binding:"required"`
}

type ChatItem struct {
	Id          int    `json:"id" db:"id"`
	Chatlist_id string `json:"chatlist_id" db:"chatlist_id"`
	User_id     string `json:"user_id" db:"user_id"`
	Username    string `json:"username" db:"username" binding:"required"`
	Description string `json:"description" db:"description"`
}

type Client struct {
	Conn     *websocket.Conn
	Message  chan *ChatItem
	Id       string `json:"id"`
	RoomId   string `json:"roomId"`
	Username string `json:"username"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type FindUserInput struct {
	Count     int    `json:"count" db:"count"`
	StartDay  string `json:"startday" db:"startday"`
	EndDay    string `json:"endday" db:"endday"`
	StartTime string `json:"starttime" db:"starttime"`
	EndTime   string `json:"endtime" db:"endtime"`
}

type UpdateListInput struct {
	Title *string `json:"title"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Description *string `json:"description"`
}

func (i UpdateItemInput) Validate() error {
	if i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UsersForChat struct {
	UsersId []int `json:"users_id"`
}

type UpdateChat struct {
	ChatName *string `json:"chat_name" db:"chatName"`
}

type UsersInfo struct {
	Id     int `json:"id" db:"id"`
	UserId int `json:"user_id" db:"user_id"`
}
