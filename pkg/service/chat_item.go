package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type ChatItemService struct {
	repo     repository.ChatItem
	listRepo repository.ChatList
}

func NewChatItemService(repo repository.ChatItem, listRepo repository.ChatList) *ChatItemService {
	return &ChatItemService{repo: repo, listRepo: listRepo}
}
func (s *ChatItemService) Create(userId, listId int, item chat.ChatItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongsto user
		return 0, nil
	}

	return s.repo.Create(listId, item)
}

func (s *ChatItemService) GetAll(userId, listId int) ([]chat.ChatItem, error) {
	return s.repo.GetAll(userId, listId)
}
