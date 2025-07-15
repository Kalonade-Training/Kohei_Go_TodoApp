package value_object

import(
	"time"
	"errors"
)

type DueDate struct{
	value time.Time
}

func NewDueDate(value string)(*DueDate, error){
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}
	//入力された時間が過去でないかバリデーション
	if parsedTime.Before(time.Now()){
		return nil, errors.New("Due date cannot be in the past")
	}
	return &DueDate{value: parsedTime}, nil
}

func (d DueDate) Value() time.Time{
	return d.value
}