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
	return &DueDate{value : parsedTime}, nil	
}

func ValidateDueDateNotPast(d *DueDate) error {
	if d != nil && d.Value().Before(time.Now()) {
		return errors.New("due date cannot be in the past")
	}
	return nil
}

func (d DueDate) Value() time.Time{
	return d.value
}