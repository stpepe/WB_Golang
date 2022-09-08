package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/stpepe/nats-task"
)

type Show interface{

}

type Send interface{
	CreateOrder (order testapp.Order) (int, error)
	FindAllOrders () (error, map[int]testapp.Order)
}

type Repository struct{
	Show
	Send
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Send: NewOrderPostgres(db),
	}
}