package todo

import (
	"errors"
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
	"goTodoApp/domain/value-object"
)

type FindTodoByIDUseCase struct {
	todoRepo repositories.ITodoRepository
}
func NewFindTodoByIDUseCase(todoRepo repositories.ITodoRepository) *FindTodoByIDUseCase {
	return &FindTodoByIDUseCase{todoRepo: todoRepo}
}
func (uc *FindTodoByIDUseCase) Execute(todoID value_object.TodoID, userID value_object.UserID) (*entities.Todo, error) {
	if todoID.Value() == "" {
		return nil, errors.New("todo ID is required")
	}
	if userID.Value() == "" {
		return nil, errors.New("user ID is required")
	}

	todo, err := uc.todoRepo.FindTodoByID(todoID, userID)
	if err != nil {
		return nil, err
	}

	if todo.UserID().Value() != userID.Value() {
		return nil, errors.New("unauthorized access to this Todo")
	}

	return todo, nil
}