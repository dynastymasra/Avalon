package service

import (
	"context"

	"github.com/dynastymasra/avalon/config"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/avalon/domain"
	"github.com/dynastymasra/avalon/domain/repository"
)

type OrderServicer interface {
	CreateOrder(context.Context, domain.Order) (*domain.Order, error)
}

type OrderService struct {
	OrderRepository repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return OrderService{
		OrderRepository: orderRepo,
	}
}

func (o OrderService) CreateOrder(ctx context.Context, order domain.Order) (*domain.Order, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"order":          order,
	})

	order.ID = uuid.NewV4().String()

	if err := o.OrderRepository.Save(ctx, order); err != nil {
		log.WithError(err).Errorln("Failed create new order")
		return nil, err
	}

	result, err := o.OrderRepository.FindByID(ctx, order.ID)
	if err != nil {
		log.WithError(err).Errorln("Failed get order from db")
		return nil, err
	}

	return result, nil
}
