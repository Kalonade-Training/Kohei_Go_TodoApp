// usecases/todo/find_by_user_with_filters.go
package todo

import (
	"fmt"
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
	value_object "goTodoApp/domain/value-object"
)

type FindByUserIDWithFiltersUseCase struct {
	todoRepo repositories.ITodoRepository
}

func NewFindByUserIDWithFiltersUseCase(todoRepo repositories.ITodoRepository) *FindByUserIDWithFiltersUseCase {
	return &FindByUserIDWithFiltersUseCase{todoRepo: todoRepo}
}

func (uc *FindByUserIDWithFiltersUseCase) Execute(userID value_object.UserID, filters repositories.TodoFilters) ([]*entities.Todo, error) {
	if userID.Value() == "" {
		return nil, fmt.Errorf("user ID is required")
	}
	return uc.todoRepo.FindByUserIDWithFilters(userID, &filters)
}