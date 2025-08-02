package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ravishgithub/krishak/authentication"
)

func TestAddAndListLands(t *testing.T) {
	// Set JWT secret for test
	// Set the JWT secret  and config path already set in handlers_test.go
	token, err := authentication.GenerateToken("admin")
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	// Test POST /lands
	addReq := httptest.NewRequest("POST", "/lands", strings.NewReader(`{
        "size": 1.5,
        "location": "Rampur",
        "soil_type": "Loamy"
    }`))
	addReq.Header.Set("Content-Type", "application/json")
	addReq.Header.Set("Authorization", token)

	addW := httptest.NewRecorder()
	AddLandHandler(addW, addReq)

	if addW.Result().StatusCode != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", addW.Result().StatusCode)
	}

	// Test GET /lands
	listReq := httptest.NewRequest("GET", "/list_lands", nil)
	listReq.Header.Set("Authorization", token)

	listW := httptest.NewRecorder()
	ListLandsHandler(listW, listReq)

	if listW.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", listW.Result().StatusCode)
	}
}
