package handlers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestAddAndListContracts(t *testing.T) {
    landID := "test-land"
    lands[landID] = Land{ID: landID, Location: "Test Location", SoilType: "Test Soil", Size: 1.5}

    addReq := httptest.NewRequest("POST", "/contracts", strings.NewReader(`{
        "contractor_id": "test-contractor",
        "land_id": "test-land",
        "start_date": "2024-01-01",
        "end_date": "2025-01-01",
        "expected_yield": 90
    }`))
    addReq.Header.Set("Content-Type", "application/json")
    addW := httptest.NewRecorder()
    AddContractHandler(addW, addReq)

    if addW.Result().StatusCode != http.StatusCreated {
        t.Fatalf("expected 201 Created")
    }

    listReq := httptest.NewRequest("GET", "/list_contracts", nil)
    listW := httptest.NewRecorder()
    ListContractsHandler(listW, listReq)

    if listW.Result().StatusCode != http.StatusOK {
        t.Fatalf("expected 200 OK")
    }
}
