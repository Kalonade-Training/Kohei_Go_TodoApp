package common

import (
	"errors"

	"github.com/gin-gonic/gin"
	"goTodoApp/domain/value-object"
)

func GetAuthUserID(c *gin.Context) (value_object.UserID, error) {
	val, ok := c.Get("userID")
	if !ok {
		return value_object.UserID{}, errors.New("missing userID in context")
	}
	rawUserID, ok := val.(string)
	if !ok {
		return value_object.UserID{}, errors.New("invalid userID format")
	}
	return value_object.FromStringUserID(rawUserID), nil
}