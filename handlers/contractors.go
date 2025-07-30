
package handlers

import (
    "encoding/json"
    "net/http"
    "sync"
    "github.com/google/uuid"
)

type Contractor struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Contact string `json:"contact"`
    Aadhar  string `json:"aadhar,omitempty"`
}

var (
    contractors = make(map[string]Contractor)
    mu          sync.Mutex
)

func AddContractorHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var contractor Contractor
    if err := json.NewDecoder(r.Body).Decode(&contractor); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    contractor.ID = uuid.New().String()

    mu.Lock()
    contractors[contractor.ID] = contractor
    mu.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(contractor)
}

func ListContractorsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    var list []Contractor
    for _, v := range contractors {
        list = append(list, v)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(list)
}
