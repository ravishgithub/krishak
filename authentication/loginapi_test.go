package authentication

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("JWT_SECRET", "ebee1a4380a9ab9a0a84b091c1f7abcf30c3428608f122dbd91e13db134b16bc")

	// Create test config.json one level above this package
	cwd, _ := os.Getwd()
	configDir := filepath.Join(cwd, "../configs")
	os.MkdirAll(configDir, 0755)

	configPath := filepath.Join(configDir, "config.json")
	config := `{
        "login": {
            "username": "admin",
            "password": "$2a$10$TyLfPLwjMxmE5nblEZZDN.UkEoFek9k3Ronc2aeLjcdaJbT31TgT2"
        },
        "server": { "port": 8080, "hostname": "localhost" },
        "database": { "username": "admin", "password": "admin123", "name": "testdb" }
    }`

	os.WriteFile(configPath, []byte(config), 0644)
	os.Setenv("CONFIG_PATH", configPath)

	os.Exit(m.Run())
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("admin")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	if token == "" {
		t.Fatal("Expected non-empty token")
	}
}

func TestIsValidToken(t *testing.T) {
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	token, _ := GenerateToken("admin")
	if !IsValidToken(token, cfg) {
		t.Errorf("Valid token was not recognized")
	}
	if IsValidToken("invalid-token", cfg) {
		t.Errorf("Invalid token was recognized as valid")
	}
}

func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	if cfg.Login.Username != "admin" {
		t.Errorf("Expected username 'admin', got '%s'", cfg.Login.Username)
	}
}

func TestNewLoginHandler(t *testing.T) {
	handler, err := NewLoginHandler()
	if err != nil {
		t.Fatalf("Failed to create login handler: %v", err)
	}
	req := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"admin","password":"admin"}`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
}

func TestNewCheckAuthHandler(t *testing.T) {
	handler, err := NewCheckAuthHandler()
	if err != nil {
		t.Fatalf("Failed to create check auth handler: %v", err)
	}
	token, _ := GenerateToken("admin")
	req := httptest.NewRequest("GET", "/check_auth", nil)
	req.Header.Set("Authorization", token)
	rr := httptest.NewRecorder()
	handler(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
}
