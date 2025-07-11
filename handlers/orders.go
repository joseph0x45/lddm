package handlers

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"server/models"
	"server/repo"
	"time"

	"github.com/google/uuid"
	"github.com/kenshaw/escpos"
)

type OrderHandler struct {
	orders *repo.OrderRepo
}

func NewOrderHandler(orders *repo.OrderRepo) *OrderHandler {
	return &OrderHandler{orders: orders}
}

func printOrder(orderData *models.OrderData) error {
	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := escpos.New(f)

	p.Init()
	p.SetAlign("center")
	p.SetFontSize(2, 3)
	p.Write("EPICES TACO\n")
	p.SetFontSize(1, 1)
	p.Write("Tel: +22891541906 / +22879806420\n")
	p.FormfeedN(1)
	p.SetAlign("left") // Left-align for product details

	p.Write("Product          Qty   Price    Total\n")
	p.Write("-------------------------------------\n")
	p.Write("-------------------------------------\n")

	p.FormfeedN(3)
	p.End()

	w.Flush()
	return nil
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
			ProductID      string `json:"product_id"`
			ProductName    string `json:"product_name"`
			ProductVariant string `json:"product_variant"`
			Quantity       int    `json:"quantity"`
			UnitPrice      int    `json:"unit_price"`
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
		IssuedAt:          time.Now().UTC().String(),
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
			ID:             uuid.NewString(),
			OrderID:        order.ID,
			ProductID:      orderItem.ProductID,
			ProductName:    orderItem.ProductName,
			ProductVariant: orderItem.ProductVariant,
			Quantity:       orderItem.Quantity,
			UnitPrice:      orderItem.UnitPrice,
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

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	data, err := h.orders.GetOrders()
	if err != nil {
		log.Printf("[ERROR] Error while getting orders: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("[ERROR] Error while marshalling data: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *OrderHandler) Print(w http.ResponseWriter, r *http.Request) {
	payload := &struct {
		OrderID string `json:"order_id"`
	}{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Printf("[ERROR] Error while decoding request body: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	orderData, err := h.orders.GetOrderByID(payload.OrderID)
	if err != nil {
		log.Printf("[ERROR] Error while getting order data: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = printOrder(orderData)
	if err != nil {
		log.Printf("[ERROR] Error while printing order: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
