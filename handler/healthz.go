package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

type HealthzHandler struct{}

func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := model.HealthzResponse{
		Message: "OK",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}