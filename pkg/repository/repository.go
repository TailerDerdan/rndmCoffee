package repository

import (
	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user chat.User) (int, error)
	GetUser(email, password string) (chat.User, error)
	GetUserEmail(token string) (chat.User, error)
	ResetPassword(email, password string) error
	DeleteUserToken(user chat.User) error
	SetUserToken(token, email string) error
}

type Profile interface {
	CreateProfile(userId int, profile chat.Profile) (int, error)
	GetProfile(userId, profileId int) (chat.Profile, error)
	EditProfile(userId, profileId int, input chat.UpdateProfile) error
	GetProfileId(userId int) (int, error)
	CreateHobby(profId int, hobbies map[string][]chat.UserHobbyInput) ([]int, error)
	GetAllHobby(profId int) ([]chat.UserHobby, error)
	DeleteHobby(profId, hobbyId int) error
	InitAllHobbies() error
	//UploadAvatar(profileId int, directory string) error
}

type ChatList interface {
	CreateList(input chat.UsersForChat) (int, error)
	RenameChat(userId, chatId int, chat chat.UpdateChat) error
	GetAllLists(userId int) ([]chat.ChatList, error)
	GetListById(userId, listId int) (chat.ChatList, error)
	DeleteList(userId, listId int) error
	UpdateList(userId, listId int, input chat.UpdateListInput) error
	FindByTime(userId int, input chat.FindUserInput) ([]int, []chat.UsersInfo, error)
	FindThreeByHobby(list_users []int) ([]chat.UserHobby, error)
	FindTwoByHobby(list_users []int) ([]chat.UserHobby, error)
	DeleteFindUsers(input chat.UsersForChat) error
	GetUserByListId(listId int) ([]int, error)
	GetUserAvatar(users_id []int) ([]string, error)
	UpdateFindUsersTable(users_info []chat.UsersInfo, count int) error
}

type ChatItem interface {
	GetUsers(userId, listId int) ([]int, error)
	Create(listId int, item chat.ChatItem) (int, error)
	GetAll(userId, listId int) ([]chat.ChatItem, error)
	// GetById(userId, itemId int) (chat.ChatItem, error)
	// Delete(userId, itemId int) error
	// Update(userId, itemId int, input chat.UpdateItemInput) error
}

type Repository struct {
	Authorization
	Profile
	ChatList
	ChatItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Profile:       NewProfilePostgres(db),
		ChatList:      NewChatListPostgres(db),
		ChatItem:      NewChatItemPostgres(db),
	}
}
