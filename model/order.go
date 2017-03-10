package model

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

// OrderStatus type
type OrderStatus string

// ArrayString type string
type ArrayString []string

// OrderStatus const value
const (
	OrderStatusCancel  OrderStatus = "CANCEL"
	OrderStatusFinish  OrderStatus = "FINISH"
	OrderStatusProcess OrderStatus = "PROCESS"
	OrderStatusNew     OrderStatus = "NEW"
)

// Order struct
type Order struct {
	ID          string      `json:"id" sql:"type:uuid" gorm:"not null;column:id;primary_key"`
	ShopID      string      `json:"shop_id" binding:"required" gorm:"not null;column:shop_id;type:varchar(255)"`
	CustomerID  string      `json:"customer_id" binding:"required" gorm:"not null;column:customer_id;type:varchar(255)"`
	OrderStatus string      `json:"order_status" gorm:"not null;column:order_status;type:varchar(255)"`
	Products    ArrayString `json:"products" binding:"required" gorm:"not null;column:products;type:varchar[]"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"-"`
	DeletedAt   *time.Time  `json:"-"`
}

// TableName override interface
func (Order) TableName() string {
	return "avalon"
}

// Value interface
func (a ArrayString) Value() (driver.Value, error) {
	if len(a) == 0 {
		return nil, nil
	}

	var buffer bytes.Buffer

	buffer.WriteString("{")
	for i, val := range a {
		buffer.WriteString(val)
		if i != len(a)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")

	return buffer.String(), nil
}

// Scan interface
func (a *ArrayString) Scan(value interface{}) error {
	bytesValue, ok := value.([]byte)
	if !ok {
		log.WithFields(log.Fields{"file": "order.go", "package": "model"}).Error("Scan source not []byte")
		return fmt.Errorf("Scan source not []byte")
	}

	stringValue := string(bytesValue)
	stringValue = strings.Replace(stringValue, "{", "", -1)
	stringValue = strings.Replace(stringValue, "}", "", -1)

	stringArray := strings.Split(stringValue, ",")
	(*a) = ArrayString(stringArray)

	return nil
}

// CheckOrderStatus if not id criteria
func CheckOrderStatus(orderStatus string) error {
	switch strings.ToUpper(orderStatus) {
	case string(OrderStatusCancel):
		return nil
	case string(OrderStatusFinish):
		return nil
	case string(OrderStatusNew):
		return nil
	case string(OrderStatusProcess):
		return nil
	default:
		return fmt.Errorf("Order status %v not found", orderStatus)
	}
}
