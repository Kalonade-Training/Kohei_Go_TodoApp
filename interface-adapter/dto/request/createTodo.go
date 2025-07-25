package dto

type Todo struct {
	Title   string  `json:"title" binding:"required"`
	Description    *string `json:"description,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	Status *string `json:"status" binding:"required"`
}