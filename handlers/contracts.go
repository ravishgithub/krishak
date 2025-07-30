
package handlers

import (
    "encoding/json"
    "net/http"
    "sync"
    "github.com/google/uuid"
)

type Contract struct {
    ID            string  `json:"id"`
    ContractorID  string  `json:"contractor_id"`
    LandID        string  `json:"land_id"`
    StartDate     string  `json:"start_date"`   // yyyy-mm-dd
    EndDate       string  `json:"end_date"`
    ExpectedYield float64 `json:"expected_yield"` // in quintals
}

var (
    contracts = make(map[string]Contract)
    contractMu sync.Mutex
)

func AddContractHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var contract Contract
    if err := json.NewDecoder(r.Body).Decode(&contract); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    contract.ID = uuid.New().String()

    contractMu.Lock()
    contracts[contract.ID] = contract
    contractMu.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(contract)
}

func ListContractsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    contractMu.Lock()
    defer contractMu.Unlock()

    var list []Contract
    for _, v := range contracts {
        list = append(list, v)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(list)
}
