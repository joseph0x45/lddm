package handlers

import (
	"embed"
	"net/http"
	"text/template"
)

type UIHandler struct {
	templateFS *embed.FS
}

func NewUIHandler(templateFS *embed.FS) *UIHandler {
	return &UIHandler{templateFS: templateFS}
}

func (h *UIHandler) RenderProductsPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(h.templateFS, "templates/login.html"))
  tmpl.Execute(w, nil)
}
