package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ravishgithub/krishak/authentication"
	"github.com/ravishgithub/krishak/handlers"
	"github.com/rs/cors"
)

func main() {
	config, err := authentication.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	loginHandler, err := authentication.NewLoginHandler()
	if err != nil {
		log.Fatalf("❌ Failed to create login handler: %v", err)
	}

	checkAuthHandler, err := authentication.NewCheckAuthHandler()
	if err != nil {
		log.Fatalf("❌ Failed to create auth check handler: %v", err)
	}

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/check_auth", checkAuthHandler)
	http.HandleFunc("/contractors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListContractorsHandler(w, r)
		case http.MethodPost:
			handlers.AddContractorHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/lands", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListLandsHandler(w, r)
		case http.MethodPost:
			handlers.AddLandHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/contracts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListContractsHandler(w, r)
		case http.MethodPost:
			handlers.AddContractHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   config.CORS.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(http.DefaultServeMux)

	addr := fmt.Sprintf("%s:%d", config.Server.Hostname, config.Server.Port)
	log.Printf("🌾 Server running at http://%s\n", addr)

	if err := http.ListenAndServe(addr, corsHandler); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
