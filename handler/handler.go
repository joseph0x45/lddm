package handler

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"server/models"
	"server/store"

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

func (h *Handler) RenderOrdersPage(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) RenderStatsPage(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) RenderCartPage(w http.ResponseWriter, r *http.Request)  {}

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
