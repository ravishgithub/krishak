package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

var validUsername = "admin"
var hashedPassword, _ = bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

func init() {
	// Load the credentials from the JSON file during initialization
	loadCredentialsFromFile("credentials.json")
}

func loadCredentialsFromFile(filename string) {
	// Read the JSON file and decode its contents into the Credentials struct
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err) // Handle error reading the file
	}

	var creds Credentials // Variable to hold the credentials
	err = json.Unmarshal(fileData, &creds)
	if err != nil {
		panic(err) // Handle error decoding JSON
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Simulating a hashed password stored in a database
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(creds.Password))
	if err != nil || creds.Username != validUsername {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate a JWT token (example)
	token := "example_jwt_token"
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

func main() {
	http.HandleFunc("/login", loginHandler)

	// Endpoint for checking authentication
	//http.HandleFunc("/check_auth", checkAuthHandler)
	//http.HandleFunc("/check_auth", authentication.CheckAuthHandler)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
