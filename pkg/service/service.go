package service

import "github.com/MerBasNik/rndmCoffee/pkg/repository"

type Autorization interface {
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
	return &Service{}
}
