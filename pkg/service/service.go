package service

import (
	todo "github.com/Brotiger/todo_go"
	"github.com/Brotiger/todo_go/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {

}

type TodoIten interface {

}

type Service struct {
	Authorization
	TodoList
	TodoIten
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}