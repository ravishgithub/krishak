package authentication

import (
	"os"
	"path/filepath"
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
            "password": "$2a$10$WcXsDJG7lsNQe08iKkH2z.rP4qqEtIZePI7vjC9dvUQx7et9RQY0u"
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
