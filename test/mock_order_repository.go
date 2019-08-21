package test

import (
	"context"

	"github.com/dynastymasra/avalon/domain"
	"github.com/dynastymasra/avalon/domain/repository"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m MockOrderRepository) Save(ctx context.Context, order domain.Order) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m MockOrderRepository) Find(ctx context.Context, query domain.Order, filter *repository.Query) ([]*domain.Order, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Order), args.Error(1)
}

func (m MockOrderRepository) Update(ctx context.Context, query domain.Order, update domain.Order) error {
	args := m.Called(ctx)
	return args.Error(0)
}
