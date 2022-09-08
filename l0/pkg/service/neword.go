package service 

import (
	"github.com/stpepe/nats-task"
	"github.com/stpepe/nats-task/pkg/repository"

)
type OrderService struct {
	repo repository.Send
}

func NewOrderService(repo repository.Send) *OrderService{
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder (order testapp.Order) (int, error){
	return s.repo.CreateOrder(order)
}