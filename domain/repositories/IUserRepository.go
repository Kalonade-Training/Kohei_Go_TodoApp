package repositories

import(
	"goTodoApp/domain/entities"
) 

type IUserRepository interface {
	FindByUsername(username string)(*entities.User, error)
	Save(user *entities.User) (*entities.User, error)
}