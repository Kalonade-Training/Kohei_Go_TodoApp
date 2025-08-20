package dto

type TodoFilters struct {
	Title       *string `form:"title"`
	Description *string `form:"description"`
	DueDateFrom *string `form:"dueDate_from"`
	DueDateTo   *string `form:"dueDate_to"`
	Status   		*string `form:"status"`
}