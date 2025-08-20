// usecases/todo/delete.go
package todo

import (
	"errors"
	"goTodoApp/domain/repositories"
	value_object "goTodoApp/domain/value-object"
)
//
type DeleteTodoUseCase struct {
	todoRepo repositories.ITodoRepository
}

func NewDeleteTodoUseCase(todoRepo repositories.ITodoRepository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{todoRepo: todoRepo}
}

func (uc *DeleteTodoUseCase) Execute(todoID value_object.TodoID, userID value_object.UserID) error {
	//引数のバリデーション
	if userID.Value() == "" {
		return errors.New("userID is required")
	}
	if todoID.Value() == "" {
		return errors.New("todoID is required")
	}

	return uc.todoRepo.Delete(todoID, userID)
}