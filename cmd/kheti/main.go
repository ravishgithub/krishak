package main

import (
	"fmt"
	"log"
	"net/http"

	"krishak.tech/kheti/authentication"
)

func main() {
	loginHandler, err := authentication.NewLoginHandler()
	if err != nil {
		log.Fatal("Error loading login handler:", err)
	}

	checkAuthHandler, err := authentication.NewCheckAuthHandler()
	if err != nil {
		log.Fatal("Error loading check auth handler:", err)
	}

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/check_auth", checkAuthHandler)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
