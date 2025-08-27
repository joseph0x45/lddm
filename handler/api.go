package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) GetData(w http.ResponseWriter, r*http.Request){
	data, err := json.Marshal(h.data)
	if err != nil {
		log.Println("[ERROR]: failed to fetch data: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
