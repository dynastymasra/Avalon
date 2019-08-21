package service

import "github.com/dynastymasra/avalon/domain/repository"

type Instance struct {
	OrderServicer OrderServicer
}

func NewInstance(orderService repository.OrderRepository) Instance {
	return Instance{
		OrderServicer: NewOrderService(orderService),
	}
}
