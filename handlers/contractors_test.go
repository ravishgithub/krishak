package handlers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestAddAndListContractors(t *testing.T) {
    addReq := httptest.NewRequest("POST", "/contractors", strings.NewReader(`{
        "name": "Test Farmer",
        "contact": "9999999999",
        "aadhar": "111122223333"
    }`))
    addReq.Header.Set("Content-Type", "application/json")
    addW := httptest.NewRecorder()
    AddContractorHandler(addW, addReq)

    if addW.Result().StatusCode != http.StatusCreated {
        t.Fatalf("expected 201 Created")
    }

    listReq := httptest.NewRequest("GET", "/list_contractors", nil)
    listW := httptest.NewRecorder()
    ListContractorsHandler(listW, listReq)

    if listW.Result().StatusCode != http.StatusOK {
        t.Fatalf("expected 200 OK")
    }
}
