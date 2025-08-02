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
	CORS     CORSConfig     `json:"cors"`
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

type CORSConfig struct {
	AllowedOrigins []string `json:"allowed_origins"`
}

var jwtSecret []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Printf("Warning: JWT_SECRET is not set")
	}
	jwtSecret = []byte(secret)
}

func LoadConfig() (Config, error) {
	var config Config
	configPath := filepath.Join("configs", "config.json")

	if customPath := os.Getenv("CONFIG_PATH"); customPath != "" {
		configPath = customPath
	}

	absPath, err := filepath.Abs(configPath)
	if err != nil {
		log.Printf("error getting absolute path: %v", err)
		return config, fmt.Errorf("error getting absolute path: %w", err)
	}
	log.Printf("Loading config from: %s", absPath)

	file, err := os.Open(absPath)
	if err != nil {
		log.Printf("error opening config file: %v", err)
		return config, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Printf("error decoding config file: %v", err)
		return config, fmt.Errorf("error decoding config file: %w", err)
	}

	return config, nil
}

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func IsValidToken(tokenString string, config Config) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["username"] == config.Login.Username {
			return true
		}
	}

	return false
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
		json.NewEncoder(w).Encode(response)
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
