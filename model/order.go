package model

import "time"

// Order struct
type Order struct {
	ID          string     `json:"id" sql:"type:uuid" gorm:"not null;column:id;primary_key"`
	ShopID      string     `json:"shop_id" gorm:"not null;column:shop_id;type:varchar(255)"`
	CustomerID  string     `json:"customer_id" gorm:"not null;column:customer_id;type:varchar(255)"`
	OrderStatus string     `json:"order_status" gorm:"not null;column:order_status;type:varchar(255)"`
	Products    []string   `json:"products" gorm:"not null;column:products;type:varchar[]"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

// TableName override interface
func (Order) TableName() string {
	return "avalon"
}
