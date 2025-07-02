package repositories

import "goTodoApp/domain/entities"

type ITodoRepository interface {
	Save(todo *entities.Todo) error
	FindByID(id string) (*entities.Todo, error)
	FindAll() ([]*entities.Todo, error)
	Delete(id string) error
	Duplicate(id string) (*entities.Todo, error)
	Update(id string, fields map[string]interface{}) error
}