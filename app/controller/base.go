package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type BaseHandler struct {
}

func (h *BaseHandler) Response(w http.ResponseWriter, status int, res any) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
	}
}
