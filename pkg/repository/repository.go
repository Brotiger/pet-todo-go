package repository

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

func NewRepositor() *Repository {
	return &Repository{}
}