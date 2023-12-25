package main

import (
	"fmt"
	"log"
	"net/http"

	"krishak.tech/kheti/authentication"
)

func main() {
	fmt.Println("Hello, Go Project!")
	http.HandleFunc("/login", authentication.LoginHandler)

	// Endpoint for checking authentication
	http.HandleFunc("/check_auth", authentication.CheckAuthHandler)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/check_auth", authentication.CheckAuthHandler)
	// Add your application logic here
	// For example, start your server or perform other initializations
}
