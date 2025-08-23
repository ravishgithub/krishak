package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddLandHandler_Success(t *testing.T) {
	land := Land{Village: "Rampur", Khasra: "K123", Acre: 2.5}
	body, _ := json.Marshal(land)
	req := httptest.NewRequest("POST", "/lands", bytes.NewReader(body))
	req.Header.Set("Authorization", ValidToken())
	rr := httptest.NewRecorder()

	AddLandHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rr.Code)
	}
	var resp Land
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("invalid response: %v", err)
	}
	if resp.Village != land.Village || resp.Khasra != land.Khasra || resp.Acre != land.Acre {
		t.Errorf("land fields not saved correctly")
	}
}

func TestAddLandHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/lands", nil)
	rr := httptest.NewRecorder()
	AddLandHandler(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}

func TestAddLandHandler_InvalidToken(t *testing.T) {
	land := Land{Village: "Rampur", Khasra: "K123", Acre: 2.5}
	body, _ := json.Marshal(land)
	req := httptest.NewRequest("POST", "/lands", bytes.NewReader(body))
	req.Header.Set("Authorization", "invalid-token")
	rr := httptest.NewRecorder()
	AddLandHandler(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rr.Code)
	}
}

func TestAddLandHandler_InvalidBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/lands", strings.NewReader("{invalid"))
	req.Header.Set("Authorization", ValidToken())
	rr := httptest.NewRecorder()
	AddLandHandler(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rr.Code)
	}
}

func TestListLandsHandler_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/lands", nil)
	req.Header.Set("Authorization", ValidToken())
	rr := httptest.NewRecorder()
	ListLandsHandler(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestListLandsHandler_Unauthorized(t *testing.T) {
	req := httptest.NewRequest("GET", "/lands", nil)
	req.Header.Set("Authorization", "invalid-token")
	rr := httptest.NewRecorder()
	ListLandsHandler(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rr.Code)
	}
}

func TestListLandsHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest("POST", "/lands", nil)
	rr := httptest.NewRecorder()
	ListLandsHandler(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}
