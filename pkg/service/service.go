package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type Autorization interface {
	CreateUser(user chat.User) (int, error) 
}

type ChatList interface {
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
	}
}
