package value_object

import (
	"errors"
)

type Description struct{
	value string
}

func NewDescription(value string)(*Description, error){
	if len(value) > 100 {
		return nil, errors.New("Description must be less than 100 characters")
	}
	return &Description{value: value}, nil
}

func (d Description) Value() string{
	return d.value
}