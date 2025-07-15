// usecases/todo/update.go
package todo

import (
	"errors"
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
	value_object "goTodoApp/domain/value-object"
)

type UpdateTodoUseCase struct {
	todoRepo repositories.ITodoRepository
}

// NewUpdateUseCase はUpdateUseCase構造体の新しいインスタンスを返す
func NewUpdateTodoUseCase(todoRepo repositories.ITodoRepository) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{todoRepo}
}

type UpdateTodoInput struct {
	Title       *value_object.Title
	Description *value_object.Description
	DueDate     *value_object.DueDate
	CompletedAt *value_object.CompletedAt
}

// Execute はTodo更新のビジネスロジックを実行
func (uc *UpdateTodoUseCase) Execute(todoID value_object.TodoID, authUserID value_object.UserID, input UpdateTodoInput) (*entities.Todo, error) {
	// IDで指定されたTodoエンティティを取得
	todo, err := uc.todoRepo.FindTodoByID(
		todoID, 
		authUserID,
	)
	if err != nil {
		return nil, errors.New("todo not found: " + err.Error())
	}
	// 作成者かどうかを確認
	if todo.UserID().Value() != authUserID.Value() {
		return nil, errors.New("unauthorized: user does not own this todo")
	}

	// 指定された値で更新
	if input.Title != nil {
		todo.UpdateTitle(input.Title.Value())
	}

	if input.Description == nil {
    todo.ClearDescription() // null（または空文字→controllerでnilに変換済）ならクリア
	} else {
    todo.UpdateDescription(input.Description.Value())
	}

	if input.DueDate == nil {
    todo.ClearDueDate()
	} else {
    todo.UpdateDueDate(input.DueDate.Value().Format("2006-01-02"))
	}

	if input.CompletedAt != nil {
		todo.MarkCompleted(input.CompletedAt.Value())
	} else {
		todo.UnmarkCompleted()
	}

	// Repository経由で保存
	updated, err := uc.todoRepo.Update(todo)
	if err != nil {
		return nil, errors.New("failed to update todo: " + err.Error())
	}

	return updated, nil
}
