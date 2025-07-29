package value_object

import (
	"errors"
	"log"
	"strings"
)

type Title struct {
	value string 
}

const TitleMaxLength = 50

//バリデーション&タイトルの作成
func NewTitle(value string) (Title, error) {
	value = strings.TrimSpace(value)
	log.Printf("[DEBUG] NewTitle called with: '%s'", value)
	if len(value) == 0 {
		return Title{}, errors.New("Title can not be empty")
	}
	if len(value) > TitleMaxLength {
		return Title{}, errors.New("Tittle is too long")
	}
	return Title{value: value}, nil 
}

//valueの値を参照するgetterメソッド
func (t Title) Value() string{
	return t.value
}

func RestoreTitle(value string) Title {
	return Title{value: value}
}