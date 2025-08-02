package authentication

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
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

type LoginConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

var jwtSecret []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("Warning: JWT_SECRET is not set")
	}
	jwtSecret = []byte(secret)
}

func LoadConfig() (Config, error) {
	var config Config
	configPath := filepath.Join("configs", "config.json")
	log.Println("filepath CONFIG_PATH:", configPath)

	if customPath := os.Getenv("CONFIG_PATH"); customPath != "" {
		configPath = customPath
		log.Println("inside  CONFIG_PATH:", configPath)
	}

	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return config, fmt.Errorf("error getting absolute path: %w", err)
	}
	log.Println("Absolute Path:", absPath)

	file, err := os.Open(absPath)
	if err != nil {
		return config, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("error decoding config file: %w", err)
	}

	return config, nil
}

func NewLoginHandler() (http.HandlerFunc, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var creds Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(config.Login.Password), []byte(creds.Password))
		if err != nil || creds.Username != config.Login.Username {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Generate a new token for the user
		token, err := GenerateToken(creds.Username)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		expiration := time.Now().Add(24 * time.Hour)

		response := map[string]interface{}{
			"token":      token,
			"expires_at": expiration.Format(time.RFC3339),
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}, nil
}

func NewCheckAuthHandler() (http.HandlerFunc, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if IsValidToken(token, config) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("User is authenticated"))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("User is not authenticated"))
		}
	}, nil
}

func IsValidToken(tokenString string, config Config) bool {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token's signing method is as expected (HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil // jwtSecret is the secret used to sign the token
	})

	if err != nil {
		// Log or handle the error accordingly
		log.Println("Error parsing token:", err)
		return false
	}

	// Check if the token is valid and if the claims match what you expect
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Here you can check the claims (e.g., username, exp) as needed
		if claims["username"] == config.Login.Username {
			return true
		}
	}

	return false
}

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
