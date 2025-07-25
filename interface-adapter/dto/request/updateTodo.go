package dto

type UpdateTodoRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	DueDate     *string `json:"due_date"`
	Status      *string `json:"status"`
	CompletedAt *string `json:"completed_at"`
}