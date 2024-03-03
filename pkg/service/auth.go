package service

import (
	"crypto/sha1"
	"fmt"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
)

const salt = "asufh9rjr23bjbb2r28qsq"

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user chat.User) (int, error) {
	user.Password = genetatePassordHash(user.Password)

	return s.repo.CreateUser(user)
}

func genetatePassordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
