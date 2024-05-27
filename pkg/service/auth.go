package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/chizhavko/todolist"
	"github.com/chizhavko/todolist/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "sdfnmjk23432knj"
	signingKey = "sdfsdfklsdmf"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	r repository.Authorization
}

func NewAuthService(r repository.Authorization) *AuthService {
	return &AuthService{r: r}
}

func (s *AuthService) CreateUser(user todolist.User) (int, error) {
	user.Password = generatePasswordWith(user.Password)
	return s.r.CreateUser(user)
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := s.r.GetUser(username, generatePasswordWith(password))
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

func (s *AuthService) ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("invalid claims type")
	}

	return claims.UserId, nil
}

func generatePasswordWith(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
