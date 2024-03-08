package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type Autorization interface {
	CreateUser(user chat.User) (int, error)
	GenerateToken(username, passowrd string) (string, error)
	ParseToken(token string) (int, error)
}

type ChatList interface {
	Create(userId int, list chat.ChatList) (int, error)
	GetAll(userId int) ([]chat.ChatList, error)
	GetById(userId, listId int) (chat.ChatList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input chat.UpdateListInput) error
}

type ChatItem interface {
	Create(userId, listId int, item chat.ChatItem) (int, error)
	GetAll(userId, listId int) ([]chat.ChatItem, error)
}

type Service struct {
	Autorization
	ChatItem
	ChatList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		ChatList:     NewChatListService(repos.ChatList),
	}
}
