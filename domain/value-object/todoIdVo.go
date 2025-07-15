package value_object

import (
	"github.com/google/uuid"
)

type TodoID struct {
	value string
}

func NewTodoID() TodoID {
	return TodoID{value: uuid.New().String()}
}

func FromStringTodoID(value string) (TodoID, error) {
	_, err := uuid.Parse(value)
	if err != nil {
		return TodoID{}, err
	}
	return TodoID{value}, nil
}

func (id TodoID) Value() string{
	return id.value
}