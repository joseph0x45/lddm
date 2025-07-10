package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/models"
	"server/repo"

	"github.com/google/uuid"
)

type ProductHandler struct {
	repo *repo.ProductRepo
}

func NewProductHandler(repo *repo.ProductRepo) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	payload := &struct {
		Name           string `json:"name"`
		Price          int    `json:"price"`
		Image          string `json:"image"`
		Description    string `json:"description"`
	}{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Printf("[ERROR] Error while decoding request body: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	product := &models.Product{
		ID:             uuid.NewString(),
		Name:           payload.Name,
		Price:          payload.Price,
		Image:          payload.Image,
		Description:    payload.Description,
	}
	err = h.repo.InsertProduct(product)
	if err != nil {
		log.Printf("[ERROR] Error while inserting new product in database: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.GetAllProducts()
	if err != nil {
		log.Printf("[ERROR] Errror while fetching all products in database: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(map[string]any{
		"products": data,
	})
	if err != nil {
		log.Printf("[ERROR] Errror while fetching all products in database: Error while marshalling data: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
}
