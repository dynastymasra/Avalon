package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type LineItems struct {
	Items []Item `json:"items" gorm:"type:jsonb;column:items;not null" validate:"required,dive"`
}

type Item struct {
	ID    string  `json:"id"  validate:"required,uuid"`
	Name  string  `json:"name"  validate:"required"`
	Price float32 `json:"price" validate:"required"`
}

func (l LineItems) Value() (driver.Value, error) {
	value, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (l *LineItems) Scan(value interface{}) error {
	source, ok := value.([]byte)
	if !ok {
		return errors.New("casting data failed")
	}

	if err := json.Unmarshal(source, &l); err != nil {
		return err
	}

	return nil
}
