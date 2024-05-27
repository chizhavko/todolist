package repository

import (
	"fmt"

	"github.com/chizhavko/todolist"
	"github.com/jmoiron/sqlx"
)

type TodoItem interface {
	Create(listId int, input todolist.TodoItem) (int, error)
	AllItems(userId, listId int) ([]todolist.TodoItem, error)
}

type TodoItemRepository struct {
	db *sqlx.DB
}

func NewTodoItemRepository(db *sqlx.DB) *TodoItemRepository {
	return &TodoItemRepository{db: db}
}

func (r *TodoItemRepository) AllItems(userId, listId int) ([]todolist.TodoItem, error) {
	var list []todolist.TodoItem

	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id=ti.id INNER JOIN %s ul on ul.list_id=li.list_id WHERE li.list_id=$1 AND ul.user_id=$2",
		todoItemsTable,
		listsItemsTable,
		usersListTable)

	err := r.db.Select(&list, query, listId, userId)
	return list, err
}

func (r *TodoItemRepository) Create(listId int, input todolist.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)

	row := tx.QueryRow(createItemQuery, input.Title, input.Description)

	err = row.Scan(&itemId)

	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	return itemId, nil
}
