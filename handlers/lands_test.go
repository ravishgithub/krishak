package handlers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestAddAndListLands(t *testing.T) {
    addReq := httptest.NewRequest("POST", "/lands", strings.NewReader(`{
        "size": 2.5,
        "location": "Test Field",
        "soil_type": "Sandy"
    }`))
    addReq.Header.Set("Content-Type", "application/json")
    addW := httptest.NewRecorder()
    AddLandHandler(addW, addReq)

    if addW.Result().StatusCode != http.StatusCreated {
        t.Fatalf("expected 201 Created")
    }

    listReq := httptest.NewRequest("GET", "/list_lands", nil)
    listW := httptest.NewRecorder()
    ListLandsHandler(listW, listReq)

    if listW.Result().StatusCode != http.StatusOK {
        t.Fatalf("expected 200 OK")
    }
}
