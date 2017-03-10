package model

import (
	"fmt"
	"strings"
	"time"
)

// OrderStatus type
type OrderStatus string

// OrderStatus const value
const (
	OrderStatusCancel  OrderStatus = "CANCEL"
	OrderStatusFinish  OrderStatus = "FINISH"
	OrderStatusProcess OrderStatus = "PROCESS"
	OrderStatusNew     OrderStatus = "NEW"
)

// Order struct
type Order struct {
	ID          string     `json:"id" sql:"type:uuid" gorm:"not null;column:id;primary_key"`
	ShopID      string     `json:"shop_id" binding:"required" gorm:"not null;column:shop_id;type:varchar(255)"`
	CustomerID  string     `json:"customer_id" binding:"required" gorm:"not null;column:customer_id;type:varchar(255)"`
	OrderStatus string     `json:"order_status" gorm:"not null;column:order_status;type:varchar(255)"`
	Products    []string   `json:"products" binding:"required" gorm:"not null;column:products;type:varchar[]"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

// TableName override interface
func (Order) TableName() string {
	return "avalon"
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
