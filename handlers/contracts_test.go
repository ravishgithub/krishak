package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ravishgithub/krishak/authentication"
)

func TestAddAndListContracts(t *testing.T) {
	// Set the JWT secret  and config path already set in handlers_test.go

	token, err := authentication.GenerateToken("admin")
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	// Test POST /contracts
	addReq := httptest.NewRequest("POST", "/contracts", strings.NewReader(`{
        "contractor_id": "contractor-123",
        "land_id": "land-456",
        "start_date": "2025-01-01",
        "end_date": "2025-12-31",
        "expected_yield": 150.5
    }`))
	addReq.Header.Set("Content-Type", "application/json")
	addReq.Header.Set("Authorization", token)

	addW := httptest.NewRecorder()
	AddContractHandler(addW, addReq)

	if addW.Result().StatusCode != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", addW.Result().StatusCode)
	}

	// Test GET /contracts
	listReq := httptest.NewRequest("GET", "/list_contracts", nil)
	listReq.Header.Set("Authorization", token)

	listW := httptest.NewRecorder()
	ListContractsHandler(listW, listReq)

	if listW.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", listW.Result().StatusCode)
	}
}
