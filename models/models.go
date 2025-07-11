package models

type Product struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Price       int    `json:"price" db:"price"`
	Image       string `json:"image" db:"image"`
	Description string `json:"description" db:"description"`
}

type Order struct {
	ID                string `json:"id" db:"id"`
	IssuedAt          string `json:"issued_at" db:"issued_at"`
	CustomerName      string `json:"customer_name" db:"customer_name"`
	CustomerPhone     string `json:"customer_phone" db:"customer_phone"`
	CustomerAddress   string `json:"customer_address" db:"customer_address"`
	Discount          int    `json:"discount" db:"discount"`
	Total             int    `json:"total" db:"total"`
	TotalWithDiscount int    `json:"total_with_discount" db:"total_with_discount"`
}

type OrderItem struct {
	ID        string `json:"id" db:"id"`
	OrderID   string `json:"order_id" db:"order_id"`
	ProductID string `json:"product_id" db:"product_id"`
	Quantity  int    `json:"quantity" db:"quantity"`
	UnitPrice int    `json:"unit_price" db:"unit_price"`
}

type OrderData struct {
	ID                string      `json:"id" db:"id"`
	IssuedAt          string      `json:"issued_at" db:"issued_at"`
	CustomerName      string      `json:"customer_name" db:"customer_name"`
	CustomerPhone     string      `json:"customer_phone" db:"customer_phone"`
	CustomerAddress   string      `json:"customer_address" db:"customer_address"`
	Discount          int         `json:"discount" db:"discount"`
	Total             int         `json:"total" db:"total"`
	TotalWithDiscount int         `json:"total_with_discount" db:"total_with_discount"`
	OrderItems        []OrderItem `json:"order_items"`
}

type ProductUpdateData struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
