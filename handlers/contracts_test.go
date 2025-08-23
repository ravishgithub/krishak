package handlers

import (
	"encoding/json"
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

	var resp Contract
	if err := json.NewDecoder(addW.Body).Decode(&resp); err != nil {
		t.Errorf("invalid response: %v", err)
	}
	if resp.ContractorID != "contractor-123" || resp.LandID != "land-456" || resp.ExpectedYield != 150.5 {
		t.Errorf("contract fields not saved correctly")
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

func TestAddContractHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/contracts", nil)
	req.Header.Set("Authorization", ValidToken())
	rr := httptest.NewRecorder()
	AddContractHandler(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}

func TestAddContractHandler_InvalidToken(t *testing.T) {
	req := httptest.NewRequest("POST", "/contracts", strings.NewReader(`{
        "contractor_id": "contractor-123",
        "land_id": "land-456",
        "start_date": "2025-01-01",
        "end_date": "2025-12-31",
        "expected_yield": 150.5
    }`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "invalid-token")
	rr := httptest.NewRecorder()
	AddContractHandler(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rr.Code)
	}
}

func TestAddContractHandler_InvalidBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/contracts", strings.NewReader("{invalid"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ValidToken())
	rr := httptest.NewRecorder()
	AddContractHandler(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rr.Code)
	}
}

func TestListContractsHandler_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/contracts", nil)
	req.Header.Set("Authorization", ValidToken())
	rr := httptest.NewRecorder()
	ListContractsHandler(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestListContractsHandler_Unauthorized(t *testing.T) {
	req := httptest.NewRequest("GET", "/contracts", nil)
	req.Header.Set("Authorization", "invalid-token")
	rr := httptest.NewRecorder()
	ListContractsHandler(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rr.Code)
	}
}

func TestListContractsHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest("POST", "/contracts", nil)
	rr := httptest.NewRecorder()
	ListContractsHandler(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}
