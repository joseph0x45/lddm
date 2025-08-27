package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) FetchGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := h.conn.FetchGroups()
	if err != nil {
		log.Println("[ERROR]: failed to fetch groups: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(groups)
	if err != nil {
		log.Println("[ERROR]: failed to fetch groups: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
