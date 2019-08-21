package domain

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusCancel  OrderStatus = "CANCEL"
	OrderStatusFinish              = "FINALIZE"
	OrderStatusProcess             = "PROCESS"
	OrderStatusNew                 = "NEW"
)

type Order struct {
	ID          string      `json:"id" gorm:"not null;column:id;primary_key" validate:"omitempty,uuid"`
	ShopID      string      `json:"shop_id" gorm:"not null;column:shop_id" validate:"required,uuid"`
	CustomerID  string      `json:"customer_id" gorm:"not null;column:customer_id" validate:"omitempty,uuid"`
	OrderStatus OrderStatus `json:"order_status" gorm:"not null;column:order_status" validate:"required"`
	LineItems
	CreatedAt time.Time  `gorm:"column:created_at;not null" json:"created_at" validate:"omitempty"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null" json:"updated_at" validate:"omitempty"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty" schema:"-"  valid:"-"`
}

func (Order) TableName() string {
	return "orders"
}
