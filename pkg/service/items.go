package service

import (
	"github.com/chizhavko/todolist"
	"github.com/chizhavko/todolist/pkg/repository"
)

type TodoItem interface {
	Create(userId, listId int, input todolist.TodoItem) (int, error)
	AllItems(userId, listId int) ([]todolist.TodoItem, error)
}

type ItemsListService struct {
	r        repository.TodoItem
	listRepo repository.TodoList
}

func NewItemsListService(r repository.TodoItem, listRepo repository.TodoList) *ItemsListService {
	return &ItemsListService{r: r, listRepo: listRepo}
}

func (s *ItemsListService) Create(userId, listId int, input todolist.TodoItem) (int, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.r.Create(listId, input)
}

func (s *ItemsListService) AllItems(userId, listId int) ([]todolist.TodoItem, error) {
	return s.r.AllItems(userId, listId)
}
