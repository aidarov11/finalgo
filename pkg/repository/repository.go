package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
}

// constructor
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
