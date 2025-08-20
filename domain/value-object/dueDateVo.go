package value_object

import (
    "errors"
    "time"
)

type DueDate struct {
    value time.Time
}

func NewDueDateFrom(value string) (*DueDate, error) {
    layout := "2006-01-02"
    parsedTime, err := time.Parse(layout, value)
    if err != nil {
        return nil, err
    }
    return &DueDate{value: parsedTime}, nil
}

func NewDueDateTo(value string) (*DueDate, error) {
    layout := "2006-01-02"
    parsedTime, err := time.Parse(layout, value)
    if err != nil {
        return nil, err
    }
    // 終日の 23:59:59 に補正
    endOfDay := time.Date(
        parsedTime.Year(),
        parsedTime.Month(),
        parsedTime.Day(),
        23, 59, 59, 0,
        parsedTime.Location(),
    )
    return &DueDate{value: endOfDay}, nil
}

func ValidateDueDateNotPast(d *DueDate) error {
    if d == nil {
        return nil
    }
    today := time.Now().Truncate(24 * time.Hour) // 今日の 00:00:00
    if d.Value().Before(today) {
        return errors.New("due date cannot be in the past")
    }
    return nil
}

func (d DueDate) Value() time.Time {
    return d.value
}