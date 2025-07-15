package value_object
import (
	"errors"
)

type RawPassword struct {
	value string
}

func NewRawPassword(value string)(RawPassword, error){
	if len(value) < 5 {
		return RawPassword{}, errors.New("Password must be at least 6 characters")
	}
	return RawPassword{value: value}, nil
}

func (u RawPassword)Value() string{
	return u.value
}