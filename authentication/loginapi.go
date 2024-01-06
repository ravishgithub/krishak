package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
type LoginConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	Login    LoginConfig    `json:"login"`
}

type ServerConfig struct {
	Port     int    `json:"port"`
	Hostname string `json:"hostname"`
}

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func loadConfig() (Config, error) {
	var config Config

	// Get the absolute path to the config.json file
	configPath := filepath.Join("configs", "config.json")
	absPath, _ := filepath.Abs(configPath)
	fmt.Println("Absolute Path:", absPath)

	// Open and read the configuration file
	file, err := os.Open(configPath)
	if err != nil {
		return config, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	// Decode JSON into a Config struct
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("error decoding config file: %w", err)
	}

	return config, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	config, errConfig := loadConfig()
	if errConfig != nil {
		fmt.Println("Error loading config:", errConfig)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	errCreds := json.NewDecoder(r.Body).Decode(&creds)
	if errCreds != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var hashedPassword, _ = bcrypt.GenerateFromPassword([]byte(config.Login.Password), bcrypt.DefaultCost)
	// Simulating a hashed password stored in a database
	errCreds = bcrypt.CompareHashAndPassword(hashedPassword, []byte(creds.Password))
	if errCreds != nil || creds.Username != config.Login.Username {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate a JWT token (example)
	token := "example_token"
	expiration := time.Now().Add(24 * time.Hour) // Set token expiration

	// Respond with the token
	response := map[string]interface{}{
		"token":      token,
		"expires_at": expiration.Format(time.RFC3339),
	}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
