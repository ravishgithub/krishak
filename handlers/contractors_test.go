package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddContractorHandler_Success(t *testing.T) {
	contractor := Contractor{Name: "Ravi Contractor", Contact: "1234567890"}
	body, _ := json.Marshal(contractor)
	req := httptest.NewRequest("POST", "/contractors", bytes.NewReader(body))
	req.Header.Set("Authorization", ValidToken())
	rr := httptest.NewRecorder()

	AddContractorHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rr.Code)
	}
	var resp Contractor
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("invalid response: %v", err)
	}
	if resp.Name != contractor.Name || resp.Contact != contractor.Contact {
		t.Errorf("contractor fields not saved correctly")
	}
}

func TestAddContractorHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/contractors", nil)
	rr := httptest.NewRecorder()
	AddContractorHandler(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}

func TestAddContractorHandler_InvalidToken(t *testing.T) {
	contractor := Contractor{Name: "Ravi Contractor", Contact: "1234567890"}
	body, _ := json.Marshal(contractor)
	req := httptest.NewRequest("POST", "/contractors", bytes.NewReader(body))
	req.Header.Set("Authorization", "invalid-token")
	rr := httptest.NewRecorder()
	AddContractorHandler(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rr.Code)
	}
}

func TestAddContractorHandler_InvalidBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/contractors", strings.NewReader("{invalid"))
	req.Header.Set("Authorization", ValidToken())
	rr := httptest.NewRecorder()
	AddContractorHandler(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rr.Code)
	}
}

func TestListContractorsHandler_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/contractors", nil)
	req.Header.Set("Authorization", ValidToken())
	rr := httptest.NewRecorder()
	ListContractorsHandler(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestListContractorsHandler_Unauthorized(t *testing.T) {
	req := httptest.NewRequest("GET", "/contractors", nil)
	req.Header.Set("Authorization", "invalid-token")
	rr := httptest.NewRecorder()
	ListContractorsHandler(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rr.Code)
	}
}

func TestListContractorsHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest("POST", "/contractors", nil)
	rr := httptest.NewRecorder()
	ListContractorsHandler(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}
