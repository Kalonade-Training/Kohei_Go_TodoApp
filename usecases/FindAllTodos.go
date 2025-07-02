package usecases

import (
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
)

type FindAllTodos struct {
	todoRepo repositories.ITodoRepository
}

func NewFindAllTodos(todoRepo repositories.ITodoRepository) *FindAllTodos {
	return &FindAllTodos{todoRepo: todoRepo}
}

func (uc *FindAllTodos) Execute() ([]*entities.Todo, error) {
	return uc.todoRepo.FindAll()
}