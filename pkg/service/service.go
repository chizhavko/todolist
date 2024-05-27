package service

import (
	"github.com/chizhavko/todolist"
	"github.com/chizhavko/todolist/pkg/repository"
)

type Authorization interface {
	CreateUser(user todolist.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Services struct {
	Authorization
	TodoList
	TodoItem
}

func NewServices(r *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(r.Authorization),
		TodoList:      NewTodoListService(r.TodoList),
		TodoItem:      NewItemsListService(r.TodoItem, r.TodoList),
	}
}
