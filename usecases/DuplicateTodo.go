package usecases

import (
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
)

type DuplicateTodo struct {
	todoRepo repositories.ITodoRepository
}

func NewDuplicateTodo(todoRepo repositories.ITodoRepository) *DuplicateTodo {
	return &DuplicateTodo{todoRepo: todoRepo}
}

func (uc *DuplicateTodo) Execute(id string) (*entities.Todo, error) {
	original, err := uc.todoRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	newTodo := original.Duplicate()
	if err := uc.todoRepo.Save(newTodo); err != nil {
		return nil, err
	}

	return newTodo, nil
}