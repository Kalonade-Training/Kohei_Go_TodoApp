package usecases

import (
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
)

type UpdateTodo struct {
	todoRepo repositories.ITodoRepository
}

func NewUpdateTodo(todoRepo repositories.ITodoRepository) *UpdateTodo {
	return &UpdateTodo{todoRepo: todoRepo}
}

func (uc *UpdateTodo) Execute(id string, fields map[string]interface{}) (*entities.Todo, error) {
	//現在の値を取得
	todo, err := uc.todoRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	//更新するデータをフィールドにマージ
	if err := todo.Update(fields); err != nil {
		return nil, err
	}

	
	// 更新処理をリポジトリに依頼
	if err := uc.todoRepo.Update(id, fields); err != nil {
		return nil, err
	}
	return todo, nil
}