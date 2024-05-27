package repository

import (
	"github.com/chizhavko/todolist"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todolist.User) (int, error)
	GetUser(username, password string) (todolist.User, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemRepository(db),
	}
}
