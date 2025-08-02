package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/ravishgithub/krishak/authentication"
)

type Land struct {
	ID       string  `json:"id"`
	Size     float64 `json:"size"` // in acres
	Location string  `json:"location"`
	SoilType string  `json:"soil_type"`
}

var (
	lands  = make(map[string]Land)
	landMu sync.Mutex
)

func AddLandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	config, err := authentication.LoadConfig()
	if err != nil {
		http.Error(w, "Internal config error", http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")
	if !authentication.IsValidToken(token, config) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var land Land
	if err := json.NewDecoder(r.Body).Decode(&land); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	land.ID = uuid.New().String()

	landMu.Lock()
	lands[land.ID] = land
	landMu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(land)
}

func ListLandsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	config, err := authentication.LoadConfig()
	if err != nil {
		http.Error(w, "Internal config error", http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")
	if !authentication.IsValidToken(token, config) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	landMu.Lock()
	defer landMu.Unlock()

	var list []Land
	for _, v := range lands {
		list = append(list, v)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}
