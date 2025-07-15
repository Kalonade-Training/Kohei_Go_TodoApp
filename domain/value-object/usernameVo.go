package value_object

import (
	"errors"
	"strings"
	"fmt"
)

type Username struct {
	value string 
}

func NewUsername(value string) (Username, error) {
	fmt.Printf("[DEBUG] Validating username: '%s'\n", value)
	if len(strings.TrimSpace(value)) == 0 {
		return Username{}, errors.New("Username cannot be empty")
	}
	return Username{value}, nil
}

func FromStringUsername(username string) (Username, error) {
	if username == "" {
		return Username{}, errors.New("username cannot be empty")
	}
	return Username{value: username}, nil
}
func (u Username) Value() string {
	return u.value
}