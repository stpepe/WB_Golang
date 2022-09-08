package service

import (
	"github.com/stpepe/nats-task/pkg/repository"
	"github.com/stpepe/nats-task"
)

type Show interface{

}

type Send interface{
	CreateOrder(order testapp.Order)(int, error)
}

type Service struct{
	Show
	Send
}

func NewService(repos *repository.Repository) *Service{
	return &Service{
		Send: NewOrderService(repos.Send),
	}
}