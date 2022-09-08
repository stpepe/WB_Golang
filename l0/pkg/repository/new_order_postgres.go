package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stpepe/nats-task"
	"encoding/json"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres (db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) CreateOrder (order testapp.Order) (int, error){
	var id int
	if err := order.Validate(); err != nil {
		return 0, err
	}
	query := fmt.Sprintf("INSERT INTO %s (\"order\") values ($1) RETURNING id", ordersTable)
	obj_order, _ := json.Marshal(order)
	row := r.db.QueryRow(query, obj_order)
	if err := row.Scan(&id); err != nil{
		return 0, err
	}
	return id, nil
}

func (r *OrderPostgres) FindAllOrders() (error, map[int]testapp.Order) {
	allOrders := make(map[int]testapp.Order)
	query := fmt.Sprintf(`SELECT "id", "order" FROM "orders"`)
	row, err := r.db.Query(query)
	for row.Next() {
		var exemp testapp.CacheOrder
		var helper testapp.Order
		err = row.Scan(&exemp.Id, &exemp.Order)
		if err != nil {
			return err, nil
		}
		un := []byte(string(exemp.Order))
		err = json.Unmarshal(un, &helper)
		allOrders[exemp.Id] = helper
	}
	if err != nil {
		return err, nil
	}
	return nil, allOrders
}

