package repository

import (
	todo "github.com/Brotiger/todo_go"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {

}

type TodoIten interface {

}

type Repository struct {
	Authorization
	TodoList
	TodoIten
}

func NewRepositor(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}