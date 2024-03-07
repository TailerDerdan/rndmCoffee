package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/MerBasNik/rndmCoffee/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "asuSDFvsrjr23bjbb2r28qsq"
	signingKey = "834hucbqp*&#)bprrqcibSGu#Rprn;"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UsedId int `json:"user_id"`
}

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

func (s *AuthService) GenerateToken(username, passowrd string) (string, error) {
	user, err := s.repo.GetUser(username, genetatePassordHash(passowrd))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
			user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func genetatePassordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
