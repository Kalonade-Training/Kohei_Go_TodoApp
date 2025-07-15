package dto

import "time"

// TodoResponse はTodo作成後のレスポンスを表現するDTOです
type TodoResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description *string    `json:"description"`
	DueDate     *string `json:"due_date"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}