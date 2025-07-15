package dto

type TodoFilters struct {
	Title       *string `form:"title"`
	Description *string `form:"description"`
	DueDateFrom *string `form:"due_date_from"`
	DueDateTo   *string `form:"due_date_to"`
	Completed   *bool   `form:"completed"`
}