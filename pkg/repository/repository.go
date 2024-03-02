package repository


type Autorization interface {
}

type ChatList interface {
}

type ChatItem interface {
}

type Repository struct {
	Autorization
	ChatItem
	ChatList
}

func NewRepository() *Repository {
	return &Repository{}
}
