package usecases

import (
	"errors"
	"goTodoApp/domain/repositories"
)

type DeleteTodo struct {
	todoRepo repositories.ITodoRepository
}

func NewDeleteTodo(todoRepo repositories.ITodoRepository) *DeleteTodo {
	return &DeleteTodo{todoRepo: todoRepo}
}

func (uc *DeleteTodo) Execute(id string) error {
	_, err := uc.todoRepo.FindByID(id)
	if err != nil {
		return errors.New("todo not found")
	}

	return uc.todoRepo.Delete(id)
}