package value_object

import (
	"github.com/lucsky/cuid"
)

type UserID struct {
	value string
}

//CUIDを使ってUserIDを新規作成
func NewUserID() UserID {
	return UserID {value: cuid.New()}
}

//既存のIDから生成
func FromStringUserID(id string) UserID {
	return UserID{value: id}
}

//値の取得
func (u UserID) Value() string {
	return u.value
}