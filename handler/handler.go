package handler

import (
	"embed"
	"html/template"
	"net/http"
	"server/db"
)

type PageData struct {
	Title string
	Data  map[string]any
}

type Handler struct {
	conn *db.DBConnection
	uiFS embed.FS
}

func NewHandler(conn *db.DBConnection, uiFS embed.FS) *Handler {
	return &Handler{
		conn: conn,
		uiFS: uiFS,
	}
}

func (h *Handler) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(h.uiFS, "ui/layouts/main.html", "ui/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := PageData{
		Title: "Home",
		Data:  nil,
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
