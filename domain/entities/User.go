package entities
import (
	"time"
	"goTodoApp/domain/value-object"
)
type User struct {
	id value_object.UserID
	username value_object.Username
	hashedPassword value_object.HashedPassword
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(
	id value_object.UserID,
	username value_object.Username,
	hashedPassword value_object.HashedPassword,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User {
		id: id,
		username: username,
		hashedPassword: hashedPassword,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
//getter
func (u *User) ID() value_object.UserID{
	return u.id
}

func (u *User) Username() value_object.Username {
	return u.username
}

func (u *User) HashedPassword() value_object.HashedPassword{
	return u.hashedPassword
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}