package repository

import (
	"context"

	"github.com/dynastymasra/avalon/domain"
	"github.com/dynastymasra/avalon/domain/repository"
	"github.com/jinzhu/gorm"
)

const (
	OrderTableName = "orders"
)

type OrderRepository struct {
	db        *gorm.DB
	TableName string
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db:        db,
		TableName: OrderTableName,
	}
}

func (o OrderRepository) Save(ctx context.Context, order domain.Order) error {
	return o.db.
		Omit("created_at").
		Table(o.TableName).
		Save(&order).Error
}

func (o OrderRepository) Find(ctx context.Context, query domain.Order, filter *repository.Query) ([]*domain.Order, error) {
	var result []*domain.Order

	db := o.db.Table(o.TableName).Where(query)
	db = translateQuery(db, filter)

	err := db.Find(&result).Error

	return result, err
}

func (o OrderRepository) FindByID(ctx context.Context, id string) (*domain.Order, error) {
	var (
		result domain.Order
		query  = domain.Order{
			ID: id,
		}
	)

	err := o.db.
		Table(o.TableName).
		Where(query).
		First(&result).Error
	return &result, err
}

func (o OrderRepository) Update(ctx context.Context, query domain.Order, update domain.Order) error {
	return o.db.
		Table(o.TableName).
		Where(query).
		Update(update).Error
}
