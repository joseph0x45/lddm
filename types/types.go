package types

import "time"

type Group struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Picture string `json:"picture" db:"picture"`
}

type Product struct {
	ID        string `json:"id" db:"id"`
	GroupID   string `json:"group_id" db:"group_id"`
	Name      string `json:"name" db:"name"`
	Variant   string `json:"variant" db:"variant"`
	Picture   string `json:"picture" db:"picture"`
	InStock   int    `json:"in_stock" db:"in_stock"`
	BasePrice int    `json:"base_price" db:"base_price"`
}

type ProductBundlePrices struct {
	ID          string `json:"id" db:"id"`
	ProductID   string `json:"product_id" db:"product_id"`
	Quantity    int    `json:"quantity" db:"quantity"`
	BundlePrice int    `json:"bundle_price" db:"bundle_price"`
}

type Order struct {
	ID                  string    `json:"id" db:"id"`
	CustomerType        string    `json:"customer_type" db:"customer_type"` //can be 'regular' or 'reseller'
	CustomerName        string    `json:"customer_name" db:"customer_name"`
	CustomerPhoneNumber string    `json:"customer_phone_number" db:"customer_phone_number"`
	CustomerAddress     string    `json:"customer_address" db:"customer_address"`
	Discount            int       `json:"discount" db:"discount"`
	SubTotal            int       `json:"subtotal" db:"subtotal"`
	Total               int       `json:"total" db:"total"`
	IssuedAt            time.Time `json:"issued_at" db:"issued_at"`
}

type OrderItem struct {
	ID               string `json:"id" db:"id"`
	OrderID          string `json:"order_id" db:"order_id"`
	ProductID        string `json:"product_id" db:"product_id"`
	Quantity         int    `json:"quantity" db:"quantity"`
	Price            int    `json:"price" db:"price"`
	UsedBundledPrice bool   `json:"used_bundled_price" db:"used_bundled_price"`
}
