package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {

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
	return &Repository{}
}