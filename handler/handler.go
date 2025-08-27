package handler

import (
	"embed"
	"html/template"
	"net/http"
	"server/db"
	"server/types"
)

type PageData struct {
	Title string
	Data  map[string]any
}

type Handler struct {
	conn *db.DBConnection
	uiFS embed.FS
	data *types.Data
}

func NewHandler(conn *db.DBConnection, uiFS embed.FS, data *types.Data) *Handler {
	return &Handler{
		conn: conn,
		uiFS: uiFS,
		data: data,
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
