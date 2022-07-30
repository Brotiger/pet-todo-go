package service

import "github.com/Brotiger/todo_go/pkg/repository"

type Authorization interface {

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
	return &Service{}
}