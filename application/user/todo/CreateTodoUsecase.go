package todo

import (
	"errors"
	"time"

	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
	"goTodoApp/domain/value-object"
)
//新しいTodoを作成して保存する処理

//Createの入力データ構造
type CreateTodoInput struct {
	UserID  value_object.UserID
	Title   value_object.Title
	Description *value_object.Description
	DueDate *value_object.DueDate
	Status value_object.Status
}

type CreateTodoUseCase struct {
	todoRepo repositories.ITodoRepository
}
//リポジトリを渡してUsecaseを初期化
func NewCreateTodoUseCase(repo repositories.ITodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{todoRepo: repo}
}
//入力を受け取り、新しいTodoを作成・保存
func (uc *CreateTodoUseCase) Execute(input CreateTodoInput) (*entities.Todo, error) {
	// Todoエンティティ生成
	todo := entities.NewTodo(
		// ID生成はサーバー側で行う
		value_object.NewTodoID(), 
		input.UserID,
		input.Title,
		input.Description,
		input.DueDate,
		// CompletedAt: 初期は nil
		nil, 
		input.Status,
		time.Now(),
		time.Now(),
	)

	// 保存
	saved, err := uc.todoRepo.Save(todo)
	if err != nil {
		return nil, errors.New("failed to save todo: " + err.Error())
	}
	return saved, nil
}