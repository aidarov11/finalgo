package service

import (
	"github.com/aidarov11/todo-app"
	"github.com/aidarov11/todo-app/pkg/repository"
)

type Authorization interface {
	// создаеть нового пользователья и возвращаеть его id также error если он есть
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
