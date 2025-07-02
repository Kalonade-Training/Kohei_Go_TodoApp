package usecases

import (
	"errors"
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
)

type FindTodoByID struct {
	todoRepo repositories.ITodoRepository
}

func NewFindTodoByID(todoRepo repositories.ITodoRepository) *FindTodoByID {
	return &FindTodoByID{todoRepo: todoRepo}
}

func (uc *FindTodoByID) Execute(id string) (*entities.Todo, error) {
	todo, err := uc.todoRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}