package value_object

import (
	"errors"
)

type Status struct {
	value string
}

// 定数（公開したい場合は大文字で）
const (
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
)

// コンストラクタ
func NewStatus(value string) (*Status, error) {
	if value != StatusInProgress && value != StatusCompleted {
		return &Status{}, errors.New("invalid status: " + value)
	}
	return &Status{value: value}, nil
}


func (s Status) Value() string {
	return s.value
}

// 比較メソッド
func (s Status) IsInProgress() bool {
	return s.value == StatusInProgress
}

func (s Status) IsCompleted() bool {
	return s.value == StatusCompleted
}