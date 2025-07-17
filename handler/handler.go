package handler

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"server/store"
)

type PageData struct {
	Title string
	Data  map[string]any
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
		Title: "Home",
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
		Title: "Products",
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
