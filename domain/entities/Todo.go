package entities

import (
	"time"
	"fmt"
	value_object "goTodoApp/domain/value-object"
)

type Todo struct {
	id          value_object.TodoID
	userID      value_object.UserID
	title       value_object.Title
	description *value_object.Description
	dueDate     *value_object.DueDate
	completedAt *value_object.CompletedAt
	status      value_object.Status
	createdAt   time.Time
	updatedAt   time.Time
}

// 新しいTodoを作成する
func NewTodo(
	id          value_object.TodoID,
	userID      value_object.UserID,
	title       value_object.Title,
	description *value_object.Description,
	dueDate     *value_object.DueDate,
	completedAt *value_object.CompletedAt,
	status      value_object.Status,
	createdAt   time.Time,
	updatedAt   time.Time,
	) *Todo {
		return &Todo{
			id:          id,
			userID:      userID,
			title:       title,
			description: description,
			dueDate:     dueDate,
			status:      status,
			createdAt:   createdAt,
			updatedAt:   updatedAt,
		}
	}
	//Todo Entityのgetter
	func (t *Todo) ID() value_object.TodoID {
		return t.id
	}

	func (t *Todo) UserID() value_object.UserID {
		return t.userID
	}

	func (t *Todo) Title() value_object.Title {
		return t.title
	}

	func (t *Todo) Description() *value_object.Description {
		return t.description
	}

	func (t *Todo) DueDate() *value_object.DueDate {
		return t.dueDate
	}

	func (t *Todo) CompletedAt() *value_object.CompletedAt {
		return t.completedAt
	}

	func (t *Todo) Status() value_object.Status {
		return t.status
	}

	func (t *Todo) UpdatedAt() time.Time{
		return t.updatedAt
	}

	func (t *Todo) CreatedAt() time.Time{
		return t.createdAt
	}

	//setter
	func (t *Todo) SetCreatedAt(ti time.Time) {
	t.createdAt = ti
	}

	func (t *Todo) SetUpdatedAt(ti time.Time) {
	t.updatedAt = ti
	}	


//title string => value-object
func (t *Todo) UpdateTitle(newTitle string) error{
	title, err := value_object.NewTitle(newTitle)
	if err != nil{
		return err
	}
	t.title = title
	t.updatedAt = time.Now()
	return nil
}

//completed_atに値が追加される時に発火
func (t *Todo) MarkCompleted(at time.Time) {
	t.completedAt = value_object.NewCompletedAt(at)
	t.updatedAt = time.Now()
}
//completed_atを空に
func (t *Todo) UnmarkCompleted() {
	t.completedAt = nil
	t.updatedAt = time.Now()
}

func (t *Todo) UpdateDescription(newDescription string) error{
	fmt.Printf("[DEBUG] UpdateDescription called with: %s\n", newDescription)
	description, err := value_object.NewDescription(newDescription)
	if err != nil {
		return err
	}
	t.description = description
	t.updatedAt = time.Now()
	return nil
}
//descriptionを空に
func (t *Todo) ClearDescription() {
	fmt.Println("[DEBUG] ClearDescription called")
	t.description = nil
	t.updatedAt = time.Now()
}

func (t *Todo) UpdateDueDate(newDueDate string) error{
	dueDate, err := value_object.NewDueDate(newDueDate)
	if err != nil {
		return err
	}
	t.dueDate = dueDate
	t.updatedAt = time.Now()
	return nil
}
func (t *Todo) ClearDueDate() {
	t.dueDate = nil
	t.updatedAt = time.Now()
}

func (t *Todo) ToCompleted(at time.Time) error {
	t.MarkCompleted(at)
	status, err := value_object.NewStatus("completed")
	if err != nil {
		return err
	}
	t.status = *status
	t.updatedAt = time.Now()
	return nil
}

func (t *Todo) ToInProgress() error {
	t.completedAt = nil
	status, err := value_object.NewStatus("in_progress")
	if err != nil {
		return err
	}
	t.status = *status
	t.updatedAt = time.Now()
	return nil
}
// // Todoを更新する
// func (t *Todo) Update(fields map[string]interface{}) error {
// 	if title, ok := fields["title"].(string); ok && title != "" {
// 		t.Title = title
// 	} else if ok && title == "" {
// 		return errors.New("title is required")
// 	}

// 	if description, ok := fields["description"].(string); ok {
// 		t.Description = description
// 	}

// 	if dueDate, ok := fields["dueDate"].(*time.Time); ok {
// 		t.DueDate = dueDate
// 	}

// 	t.UpdatedAt = time.Now()
// 	return nil
// }

// Todoを複製する
func (t *Todo) Duplicate() (*Todo, error) {
	newID := value_object.NewTodoID()
	newTitle, _ := value_object.NewTitle(t.title.Value() + "のコピー")
	newStatus, err := value_object.NewStatus("in_progress")
	if err != nil {
		return nil, err
	}
	return &Todo{
		id:          newID,
		userID:      t.userID, // 同じユーザー
		title:       newTitle,
		description: t.description, // 本文はそのまま
		dueDate:     nil, // 複製時は期限なし
		completedAt: nil, // 完了状態はリセット
		status:      *newStatus,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	},nil
}