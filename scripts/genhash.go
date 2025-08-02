//go:build ignore

package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "admin"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error generating bcrypt hash: %v", err)
	}

	log.Println("Generated bcrypt hash for password 'admin':")
	log.Println(string(hash))
}
