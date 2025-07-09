package models

import "time"

type Product struct {
	ID              string `json:"id" db:"id"`
	Name            string `json:"name" db:"name"`
	Price           int    `json:"price" db:"price"`
	WholesalePrice  int    `json:"wholesale_price" db:"wholesale_price"`
	Image           string `json:"image" db:"image"`
	QuantityInStock int    `json:"quantity_in_stock" db:"quantity_in_stock"`
	Description     string `json:"description" db:"description"`
}

type Order struct {
	ID           string    `json:"id" db:"id"`
	IssuedAt     time.Time `json:"issued_at" db:"issued_at"`
	CustomerName string    `json:"customer_name" db:"customer_name"`
	Total        int       `json:"total" db:"total"`
}

type OrderItem struct {
	ID        string `json:"id" db:"id"`
	OrderID   string `json:"order_id" db:"order_id"`
	ProductID string `json:"product_id" db:"product_id"`
	Quantity  int    `json:"quantity" db:"quantity"`
	UnitPrice int    `json:"unit_price" db:"unit_price"`
}

type ProductUpdateData struct {
	Name            string `json:"name"`
	Price           int    `json:"price"`
	WholesalePrice  int    `json:"wholesale_price"`
	Image           string `json:"image"`
	QuantityInStock int    `json:"quantity_in_stock"`
	Description     string `json:"description"`
}
