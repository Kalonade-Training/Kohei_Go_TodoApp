package dto

type UpdateTodoRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	DueDate     *string `json:"due_date"`
	CompletedAt *string `json:"completed_at"`
}