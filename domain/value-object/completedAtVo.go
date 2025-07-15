package value_object

import (
	"time"
)

type CompletedAt struct {
	value time.Time
}

func NewCompletedAt(value time.Time) *CompletedAt{
	return &CompletedAt{value: value}
}

func (c CompletedAt) Value() time.Time{
	return c.value
}