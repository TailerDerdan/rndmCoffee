package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type ChatListService struct {
	repo repository.ChatList
}

func NewChatListService(repo repository.ChatList) *ChatListService {
	return &ChatListService{repo: repo}
}

func (s *ChatListService) Create(userId int, list chat.ChatList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *ChatListService) GetAll(userId int) ([]chat.ChatList, error) {
	return s.repo.GetAll(userId)
}

func (s *ChatListService) GetById(userId, listId int) (chat.ChatList, error) {
	return s.repo.GetById(userId, listId)
}
