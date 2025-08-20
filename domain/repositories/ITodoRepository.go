package repositories

import (
	"goTodoApp/domain/entities"
	value_object "goTodoApp/domain/value-object"

)

type ITodoRepository interface {
	Save(todo *entities.Todo) (*entities.Todo,error)
	FindTodoByID(todoID value_object.TodoID, userID value_object.UserID) (*entities.Todo, error)
	FindByUserIDWithFilters(userID value_object.UserID, filters *TodoFilters) ([]*entities.Todo, error)
	Delete(todoid value_object.TodoID, userID value_object.UserID ) error
	Update(todo *entities.Todo) (*entities.Todo, error)   
	Duplicate(todoid value_object.TodoID, userID value_object.UserID ) (*entities.Todo, error)
}