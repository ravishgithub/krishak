package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ravishgithub/krishak/authentication"
)

func TestAddAndListContractors(t *testing.T) {
	// Set the JWT secret  and config path already set in handlers_test.go

	// Generate a valid token
	token, err := authentication.GenerateToken("admin")
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	// Set up POST request to /contractors
	addReq := httptest.NewRequest("POST", "/contractors", strings.NewReader(`{
        "name": "Test Farmer",
        "contact": "9999999999",
        "aadhar": "111122223333"
    }`))
	addReq.Header.Set("Content-Type", "application/json")
	addReq.Header.Set("Authorization", token)

	addW := httptest.NewRecorder()
	AddContractorHandler(addW, addReq)

	if addW.Result().StatusCode != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", addW.Result().StatusCode)
	}

	// GET request to /list_contractors
	listReq := httptest.NewRequest("GET", "/list_contractors", nil)
	listReq.Header.Set("Authorization", token)

	listW := httptest.NewRecorder()
	ListContractorsHandler(listW, listReq)

	if listW.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", listW.Result().StatusCode)
	}
}
