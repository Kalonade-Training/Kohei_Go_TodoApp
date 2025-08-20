package mapper

import (
	"goTodoApp/domain/entities"
	"goTodoApp/domain/value-object"
	"goTodoApp/infrastructures/model"
)

func EntityToUserModel(user entities.User) model.User {
	return model.User{
		ID:             user.ID().Value(),
		Username:       user.Username().Value(),
		HashedPassword: user.HashedPassword().Value(),
		CreatedAt:      user.CreatedAt(),
		UpdatedAt:      user.UpdatedAt(),
	}
}


func ModelToUserEntity(m model.User) (*entities.User,error){
	userID:= value_object.FromStringUserID(m.ID)

	username, err := value_object.NewUsername(m.Username)
	if err != nil {
		return nil, err
	}
	hashedPassword:= value_object.NewHashedPassword(m.HashedPassword)

	return entities.NewUser(
		userID,
		username,
		hashedPassword,
		m.CreatedAt,
		m.UpdatedAt,
	),nil
}