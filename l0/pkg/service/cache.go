package service

import (
	"sync"
	"github.com/stpepe/nats-task/pkg/repository"
	"github.com/stpepe/nats-task"
)

type Cache struct {
    sync.RWMutex
    orders map[int]testapp.Order
}

func NewCache() *Cache {
    orders := make(map[int]testapp.Order)
    cache := Cache{
		orders: orders,
	}
    return &cache
}

func (c *Cache) Set(key int, value testapp.Order) {
    c.Lock()
    defer c.Unlock()

    c.orders[key] = value

}

func (c *Cache) Get(key int) (testapp.Order, bool) {
    c.RLock()
    defer c.RUnlock()
	_, ok := c.orders[key]
	if ok{
		return c.orders[key], true
	}
	return c.orders[key], false
}

func (c *Cache) GetAll() (map[int]testapp.Order) {
    c.RLock()
    defer c.RUnlock()

    return c.orders
}


func (c *Cache) RecoverCache(repo repository.Send){
	err, js_obj := repo.FindAllOrders()
	if err != nil{
		return
	}
	for key, value := range js_obj {
        c.Set(key, value)
    }
	
}