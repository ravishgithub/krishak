package authentication

import (
    "net/http"
    "net/http/httptest"
    "os"
    "strings"
    "testing"
    "io"
)

func TestLoginSuccess(t *testing.T) {
    os.Setenv("JWT_SECRET", "krishakdevsupersecret")

    testConfig := `{
        "server": {"port": 8080, "hostname": "localhost"},
        "database": {"username": "admin", "password": "admin123", "name": "mydatabase"},
        "login": {"username": "admin", "password": "$2b$12$z/508EYjbHG/aZG1YnJt8eVIlePdDZpbD7DQU2wENp3kRtDfrbz2u"}
    }`
    os.MkdirAll("configs", 0o755)
    os.WriteFile("configs/config.json", []byte(testConfig), 0o644)

    handler, err := NewLoginHandler()
    if err != nil {
        t.Fatalf("could not get login handler: %v", err)
    }

    payload := `{"username":"admin","password":"krishak2024"}`
    req := httptest.NewRequest("POST", "/login", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    handler(w, req)

    res := w.Result()
    if res.StatusCode != http.StatusOK {
        t.Fatalf("expected 200 OK, got %d", res.StatusCode)
    }

    body, _ := io.ReadAll(res.Body)
    if !strings.Contains(string(body), "token") {
        t.Fatalf("response does not contain token: %s", string(body))
    }
}
