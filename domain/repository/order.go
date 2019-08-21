package repository

import (
	"context"

	"github.com/dynastymasra/avalon/domain"
)

type OrderRepository interface {
	Save(context.Context, domain.Order) error
	Find(context.Context, domain.Order, *Query) ([]*domain.Order, error)
	Update(ctx context.Context, where domain.Order, update domain.Order) error
}
