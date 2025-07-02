package usecases

import (
	"errors"
	"time"
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
)

type CreateTodo struct {
	todoRepo repositories.ITodoRepository
}

func NewCreateTodo(todoRepo repositories.ITodoRepository) *CreateTodo {
	return &CreateTodo{todoRepo: todoRepo}
}

func (uc *CreateTodo) Execute(title string, description string, dueDate *time.Time) (*entities.Todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	todo, err := entities.NewTodo(title, description, dueDate)
	if err != nil {
		return nil, err
	}

	err = uc.todoRepo.Save(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}