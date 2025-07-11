package repo

import (
	"fmt"
	"server/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) InsertOrder(order *models.Order, orderItems []models.OrderItem) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	const orderQuery = `
    insert into orders(
      id, issued_at, customer_name, customer_phone,
      customer_address, discount, total,
      total_with_discount
    )
    values(
      :id, :issued_at, :customer_name, :customer_phone,
      :customer_address, :discount, :total,
      :total_with_discount
    )
  `
	_, err = tx.NamedExec(orderQuery, order)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("Error while rolling back transaction: %w", rollbackErr)
		}
		return fmt.Errorf("Error while inserting order: %w", err)
	}
	const orderItemsQuery = `
    insert into order_items(
      id, order_id, product_id,
      quantity, unit_price
    )
    values (
      :id, :order_id, :product_id,
      :quantity, :unit_price
    )
  `
	_, err = tx.NamedExec(orderItemsQuery, orderItems)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("Error while rolling back transaction: %w", rollbackErr)
		}
		return fmt.Errorf("Error while inserting order items: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Error while commiting transaction: %w", err)
	}
	return nil
}

func (r *OrderRepo) GetOrders() ([]models.OrderData, error) {
	const getOrdersQuery = "select * from orders"
	var orders []models.OrderData
	err := r.db.Select(&orders, getOrdersQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while getting orders: %w", err)
	}
	const getOrderItemsQuery = "select * from order_items where order_id=?"
	for i, order := range orders {
		var items []models.OrderItem
		err = r.db.Select(&items, getOrderItemsQuery, order.ID)
		if err != nil {
			return nil, fmt.Errorf("Error while getting order item: %w", err)
		}
		orders[i].OrderItems = items
	}
	return orders, err
}

func (r *OrderRepo) GetOrderByID(id string) (*models.OrderData, error) {
	order := &models.OrderData{}
	const getOrderQuery = "select * from orders where id=?"
	err := r.db.Get(order, getOrderQuery, id)
	if err != nil {
		return nil, fmt.Errorf("Error while getting order by id: %w", err)
	}
	var orderItems []models.OrderItem
	const getOrderItemsQuery = "select * from order_items where order_id=?"
	err = r.db.Select(&orderItems, getOrderItemsQuery, order.ID)
	if err != nil {
		return nil, fmt.Errorf("Error while getting order items: %w", err)
	}
	order.OrderItems = orderItems
	return order, nil
}
