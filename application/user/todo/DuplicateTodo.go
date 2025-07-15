// usecases/todo/duplicate.go
package todo

import (
	"errors"
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
	value_object "goTodoApp/domain/value-object"
)

type DuplicateTodoUseCase struct {
	todoRepo repositories.ITodoRepository
}

func NewDuplicateTodoUseCase(todoRepo repositories.ITodoRepository) *DuplicateTodoUseCase {
	return &DuplicateTodoUseCase{todoRepo: todoRepo}
}

func (uc *DuplicateTodoUseCase) Execute(todoid value_object.TodoID, userid value_object.UserID) (*entities.Todo, error) {
	if todoid.Value() == "" {
		return nil, errors.New("todo ID is required")
	}
	if userid.Value() == "" {
		return nil, errors.New("user ID is required")
	}
	duplicated, err := uc.todoRepo.Duplicate(todoid, userid)
	if err != nil {
		return nil, errors.New("failed to duplicate todo: " + err.Error())
	}
	return duplicated, nil
}