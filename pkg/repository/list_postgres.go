package repository

import (
	"fmt"
	"strings"

	"github.com/chizhavko/todolist"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoList interface {
	Create(userId int, input todolist.TodoList) (int, error)
	AllItems(userId int) ([]todolist.TodoList, error)
	GetListById(userId int, listId int) (todolist.TodoList, error)
	DeleteList(userId int, listId int) error
	UpdateList(userId int, listId int, update todolist.TodoListUpdate) error
}

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (s *TodoListPostgres) UpdateList(userId int, listId int, update todolist.TodoListUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if update.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *update.Title)
		argId++
	}

	if update.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *update.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id=ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListTable, setQuery, usersListTable, argId, argId+1)

	args = append(args, listId, userId)

	logrus.Debugf("updating query: %s", query)
	logrus.Debugf("updating args: %s", args)

	_, err := s.db.Exec(query, args...)

	return err
}

func (s *TodoListPostgres) DeleteList(userId int, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id=ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", todoListTable, usersListTable)
	_, err := s.db.Exec(query, userId, listId)
	return err
}

func (s *TodoListPostgres) GetListById(userId int, listId int) (todolist.TodoList, error) {
	var list todolist.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id WHERE ul.user_id=$1 AND ul.list_id=$2", todoListTable, usersListTable)

	err := s.db.Get(&list, query, userId, listId)

	return list, err
}

func (s *TodoListPostgres) AllItems(userId int) ([]todolist.TodoList, error) {
	var list []todolist.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id WHERE ul.user_id=$1", todoListTable, usersListTable)

	err := s.db.Select(&list, query, userId)
	return list, err
}

func (s *TodoListPostgres) Create(userId int, input todolist.TodoList) (int, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2 ) RETURNING id", todoListTable)
	row := tx.QueryRow(createListQuery, input.Title, input.Description)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2) RETURNING id", usersListTable)

	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
