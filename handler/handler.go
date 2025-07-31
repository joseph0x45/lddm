package handler

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"server/models"
	"server/store"
	"time"

	"github.com/google/uuid"
)

var SIDE_BAR_ITEMS = []map[string]string{
	{
		"Text": "Home",
		"Url":  "/home",
	},
	{
		"Text": "Products",
		"Url":  "/products",
	},
	{
		"Text": "Orders",
		"Url":  "/orders",
	},
	{
		"Text": "Stats",
		"Url":  "/stats",
	},
}

type PageData struct {
	Title        string
	SidebarItems []map[string]string
	Data         map[string]any
}

type Handler struct {
	store       *store.Store
	templatesFS *embed.FS
}

func NewHandler(store *store.Store, templatesFS *embed.FS) *Handler {
	return &Handler{store: store, templatesFS: templatesFS}
}

func (h *Handler) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(h.templatesFS, "templates/layouts/main_layout.html", "templates/home.html")
	if err != nil {
		log.Printf("[ERROR] Error while parsing templates: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := PageData{
		Title:        "Home",
		SidebarItems: SIDE_BAR_ITEMS,
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

func (h *Handler) RenderProductsPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(h.templatesFS, "templates/layouts/main_layout.html", "templates/products.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products, err := h.store.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := PageData{
		Title:        "Products",
		SidebarItems: SIDE_BAR_ITEMS,
		Data: map[string]any{
			"Products": products,
		},
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) RenderOrdersPage(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) RenderStatsPage(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Name        string `json:"name"`
		Variant     string `json:"variant"`
		Price       int    `json:"price"`
		Description string `json:"description"`
		Image       string `json:"image"`
		InStock     int    `json:"in_stock"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("[ERROR]: Failed to decode JSON: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newProduct := &models.Product{
		ID:          uuid.NewString(),
		Name:        payload.Name,
		Variant:     payload.Variant,
		Price:       payload.Price,
		Description: payload.Description,
		Image:       payload.Image,
		InStock:     payload.InStock,
	}
	err = h.store.InsertProduct(newProduct)
	if err != nil {
		log.Println("[ERROR] Error while inserting product: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) SaveOrder(w http.ResponseWriter, r *http.Request) {
	type saveOrderPayload struct {
		CustomerName    string `json:"customer_name"`
		CustomerPhone   string `json:"customer_phone"`
		CustomerAddress string `json:"customer_address"`
		Discount        int    `json:"discount"`
		Total           int    `json:"total"`
		SubTotal        int    `json:"subtotal"`
		Products        []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Price    int    `json:"price"`
			Quantity int    `json:"quantity"`
			Variant  string `json:"variant"`
		}
	}
	payload := saveOrderPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("[ERROR] Error while saving order: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newOrder := models.Order{
		ID:              uuid.NewString(),
		IssuedAt:        time.Now().UTC().String(),
		CustomerName:    payload.CustomerName,
		CustomerPhone:   payload.CustomerPhone,
		CustomerAddress: payload.CustomerAddress,
		Discount:        payload.Discount,
		Total:           payload.Total,
		SubTotal:        payload.SubTotal,
	}
	orderItems := make([]models.OrderItem, 0)
	for _, product := range payload.Products {
		newOrderItem := models.OrderItem{
			ID:             uuid.NewString(),
			OrderID:        newOrder.ID,
			ProductID:      product.ID,
			ProductName:    product.Name,
			ProductVariant: product.Variant,
			Price:      product.Price,
			Quantity:       product.Quantity,
		}
		orderItems = append(orderItems, newOrderItem)
	}
	err = h.store.InsertOrder(&newOrder, orderItems)
	if err != nil {
		log.Println("[ERROR] Error while saving order: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
