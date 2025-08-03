package store

import (
	"fmt"
	"server/models"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) InsertProduct(product *models.Product) error {
	_, err := s.db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}
	const query = `
    insert into products (
      id, name, variant, price,
      image, description, in_stock
    )
    values (
      :id, :name, :variant, :price,
      :image, :description, :in_stock
    );
  `
	_, err = s.db.NamedExec(query, product)
	return err
}

func (s *Store) GetAllProducts() ([]models.Product, error) {
	products := make([]models.Product, 0)
	const query = `select * from products`
	err := s.db.Select(&products, query)
	return products, err
}

func (s *Store) InsertOrder(order *models.Order, orderItems []models.OrderItem) error {
	_, err := s.db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	const orderQuery = `
    insert into orders(
      id, issued_at, customer_name, customer_phone,
      customer_address, discount, total,
      subtotal
    )
    values(
      :id, :issued_at, :customer_name, :customer_phone,
      :customer_address, :discount, :total,
      :subtotal
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
      product_name, product_variant,
      quantity, price
    )
    values (
      :id, :order_id, :product_id,
      :product_name, :product_variant,
      :quantity, :price
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

func (s *Store) GetAllOrders() ([]models.OrderData, error) {
	const getOrdersQuery = "select * from orders"
	var orders []models.OrderData
	err := s.db.Select(&orders, getOrdersQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while getting orders: %w", err)
	}
	const getOrderItemsQuery = "select * from order_items where order_id=?"
	for i, order := range orders {
		var items []models.OrderItem
		err = s.db.Select(&items, getOrderItemsQuery, order.ID)
		if err != nil {
			return nil, fmt.Errorf("Error while getting order item: %w", err)
		}
		orders[i].OrderItems = items
	}
	return orders, err
}

func (s *Store) GetOrderByID(id string) (*models.OrderData, error) {
	order := &models.OrderData{}
	const getOrderQuery = "select * from orders where id=?"
	err := s.db.Get(order, getOrderQuery, id)
	if err != nil {
		return nil, fmt.Errorf("Error while getting order by id: %w", err)
	}
	var orderItems []models.OrderItem
	const getOrderItemsQuery = "select * from order_items where order_id=?"
	err = s.db.Select(&orderItems, getOrderItemsQuery, order.ID)
	if err != nil {
		return nil, fmt.Errorf("Error while getting order items: %w", err)
	}
	order.OrderItems = orderItems
	return order, nil
}

func (s *Store) DeleteOrderByID(id string) error {
	_, err := s.db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}
	query := `
    delete from orders where id=?
  `
	_, err = s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
