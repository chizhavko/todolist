package service

import (
	"github.com/chizhavko/todolist"
	"github.com/chizhavko/todolist/pkg/repository"
)

type TodoList interface {
	Create(userId int, input todolist.TodoList) (int, error)
	AllItems(userId int) ([]todolist.TodoList, error)
	GetListById(userId int, listId int) (todolist.TodoList, error)
	DeleteList(userId int, listId int) error
	UpdateList(userId int, listId int, update todolist.TodoListUpdate) error
}

type TodoListService struct {
	r repository.TodoList
}

func NewTodoListService(r repository.TodoList) *TodoListService {
	return &TodoListService{r: r}
}

func (s *TodoListService) Create(userId int, input todolist.TodoList) (int, error) {
	return s.r.Create(userId, input)
}

func (s *TodoListService) AllItems(userId int) ([]todolist.TodoList, error) {
	return s.r.AllItems(userId)
}

func (s *TodoListService) GetListById(userId int, listId int) (todolist.TodoList, error) {
	return s.r.GetListById(userId, listId)
}

func (s *TodoListService) DeleteList(userId int, listId int) error {
	return s.r.DeleteList(userId, listId)
}

func (s *TodoListService) UpdateList(userId int, listId int, update todolist.TodoListUpdate) error {
	return s.r.UpdateList(userId, listId, update)
}
