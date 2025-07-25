package repositories
import (
	"time"
)
//updateの際に使用
type TodoFilters struct {
	Title *string
	Description *string
	DueDateFrom *time.Time
	DueDateTo *time.Time
	Status *string
}