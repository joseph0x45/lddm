package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/models"
	"server/repo"
	"time"

	"github.com/google/uuid"
)

type OrderHandler struct {
	orders *repo.OrderRepo
}

func NewOrderHandler(orders *repo.OrderRepo) *OrderHandler {
	return &OrderHandler{orders: orders}
}

func (h *OrderHandler) SaveOrder(w http.ResponseWriter, r *http.Request) {
	payload := &struct {
		CustomerName      string `json:"customer_name"`
		CustomerPhone     string `json:"customer_phone"`
		CustomerAddress   string `json:"customer_address"`
		Discount          int    `json:"discount"`
		Total             int    `json:"total"`
		TotalWithDiscount int    `json:"total_with_discount"`
		OrderItems        []struct {
			ProductID string `json:"product_id"`
			Quantity  int    `json:"quantity"`
			UnitPrice int    `json:"unit_price"`
    } `json:"order_items"`
	}{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Printf("[ERROR] Failed to decode request body: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	order := &models.Order{
		ID:                uuid.NewString(),
		IssuedAt:          time.Now().UTC(),
		CustomerName:      payload.CustomerName,
		CustomerPhone:     payload.CustomerPhone,
		CustomerAddress:   payload.CustomerAddress,
		Discount:          payload.Discount,
		Total:             payload.Total,
		TotalWithDiscount: payload.TotalWithDiscount,
	}
	orderItems := make([]models.OrderItem, 0)
	for _, orderItem := range payload.OrderItems {
		orderItems = append(orderItems, models.OrderItem{
			ID:        uuid.NewString(),
			OrderID:   order.ID,
			ProductID: orderItem.ProductID,
			Quantity:  orderItem.Quantity,
			UnitPrice: orderItem.UnitPrice,
		})
	}
	err = h.orders.InsertOrder(order, orderItems)
	if err != nil {
		log.Printf("[ERROR] Error while inserting order: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("[INFO] Order Created with ID: %s\n", order.ID)
  w.WriteHeader(http.StatusCreated)
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {}
