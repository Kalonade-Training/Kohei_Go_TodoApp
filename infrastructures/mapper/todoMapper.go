package mapper

import(
	"goTodoApp/domain/entities"
	"goTodoApp/domain/value-object"
	"goTodoApp/infrastructures/model"
	"time"
	"strings"
)
//Entity->Model型に変換
func EntityToTodoModel(todo entities.Todo) model.Todo{
	var description *string
	if desc := todo.Description(); desc != nil {
		val := strings.TrimSpace(desc.Value())
		description = &val
	}

	var dueDate *time.Time
	if todo.DueDate() != nil {
		d := todo.DueDate().Value()
		dueDate = &d 
	} 
	var completedAt *time.Time
    if todo.CompletedAt() != nil {
        c := todo.CompletedAt().Value()
        completedAt = &c
    }

	return model.Todo{
		ID:          strings.TrimSpace(todo.ID().Value()),
		UserID:      strings.TrimSpace(todo.UserID().Value()),
		Title:       strings.TrimSpace(todo.Title().Value()),
		Description: description,
		DueDate:     dueDate,
		CompletedAt: completedAt,
		Status:      strings.TrimSpace(todo.Status().Value()),
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
}
//Model->Entity型に変換
func ModelToEntity(m model.Todo) (*entities.Todo,error) {
	//必須項目
	id, err := value_object.FromStringTodoID(m.ID)
	if err != nil {
		return nil, err
	}
	userID:= value_object.FromStringUserID(m.UserID)

	title := value_object.RestoreTitle(m.Title)

	status := value_object.RestoreStatus(m.Status)
	//任意項目
	var description *value_object.Description
	if m.Description != nil && *m.Description != "" {
		d := value_object.RestoreDescription(*m.Description)
		description = d
	}
	var dueDate *value_object.DueDate
	if m.DueDate != nil {
		d, err := value_object.NewDueDateFrom(m.DueDate.Format("2006-01-02"))
		if err != nil {
			return nil, err
		}
		dueDate = d
	}
	var completedAt *value_object.CompletedAt
	if m.CompletedAt != nil {
		c := value_object.NewCompletedAt(*m.CompletedAt)
		completedAt = c
	}
	return entities.NewTodo(
		id,
		userID,
		title,
		description,
		dueDate,
		completedAt,
		*status,
		m.CreatedAt,
		m.UpdatedAt,
	),nil
}