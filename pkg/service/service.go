package service

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

type Authorization interface {
	CreateUser(user chat.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	ForgotPassword(input string) (string, error)
	ResetPassword(resetToken, password string) error
}

type Profile interface {
	CreateProfile(userId int, profile chat.Profile) (int, error)
	GetProfile(userId, profileId int) (chat.Profile, error)
	EditProfile(userId, profileId int, input chat.UpdateProfile) error

	CreateHobby(profId int, hobbies map[string][]chat.UserHobbyInput) ([]int, error)
	GetAllHobby(profId int) ([]chat.UserHobby, error)
	DeleteHobby(profId, hobbyId int) error
	InitAllHobbies() error
	//UploadAvatar(profileId int, directory string) error
}

type ChatList interface {
	CreateList(requestCreateList chat.RequestCreateList) (int, string, error)
	GetAllLists(userId int) ([]chat.ChatList, error)
	GetListById(userId, listId int) (chat.ChatList, error)
	DeleteList(userId, listId int) error
	UpdateList(userId, listId int, input chat.UpdateListInput) error
	FindByTime(userId int, input chat.FindUserInput) (int, error)
	FindByHobby(userId1, userId2 int) ([]chat.UserHobby, error)
}

type ChatItem interface {
	GetUsers(userId, listId int) ([]int, error)
	Create(userId, listId int, item chat.ChatItem) (int, error)
	GetAll(userId, listId int) ([]chat.ChatItem, error)
	// GetById(userId, itemId int) (chat.ChatItem, error)
	// Delete(userId, itemId int) error
	// Update(userId, itemId int, input chat.UpdateItemInput) error
}

type Service struct {
	Authorization
	Profile
	ChatList
	ChatItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Profile:       NewProfileService(repos.Profile),
		ChatList:      NewChatListService(repos.ChatList),
		ChatItem:      NewChatItemService(repos.ChatItem, repos.ChatList),
	}
}
