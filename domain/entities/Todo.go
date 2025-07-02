package entities

import (
	"errors"
	"time"
	"github.com/google/uuid"
)

type Todo struct {
    ID          string     `gorm:"type:char(36);primaryKey"`          // UUIDなど固定長文字列
    Title       string     `gorm:"type:varchar(255);not null"`        // 255文字の文字列、null不可
    Description string     `gorm:"type:text"`                         // テキスト型
    DueDate     *time.Time `gorm:"type:date"`                         // 日付だけ（yyyy-mm-dd）
    CompletedAt *time.Time                                 // 省略すると自動判別
    CreatedAt   time.Time  `gorm:"autoCreateTime"`                    // 作成日時を自動セット
    UpdatedAt   time.Time  `gorm:"autoUpdateTime"`                    // 更新日時を自動セット
}

// 新しいTodoを作成する
func NewTodo(title string, description string, dueDate *time.Time) (*Todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	return &Todo{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// Todoを更新する
func (t *Todo) Update(fields map[string]interface{}) error {
	if title, ok := fields["title"].(string); ok && title != "" {
		t.Title = title
	} else if ok && title == "" {
		return errors.New("title is required")
	}

	if description, ok := fields["description"].(string); ok {
		t.Description = description
	}

	if dueDate, ok := fields["dueDate"].(*time.Time); ok {
		t.DueDate = dueDate
	}

	t.UpdatedAt = time.Now()
	return nil
}

// Todoを複製する
func (t *Todo) Duplicate() *Todo {
	return &Todo{
		ID: uuid.New().String(),
		Title:       t.Title + "のコピー",
		Description: t.Description,
		DueDate:     nil,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}