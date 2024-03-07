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
}

type ChatItem interface {
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
