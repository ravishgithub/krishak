
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/ravishgithub/krishak/handlers"
    "github.com/ravishgithub/krishak/authentication"

)

func withAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")

        config, err := authentication.LoadConfig()
        if err != nil || !authentication.IsValidToken(token, config) {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("Unauthorized"))
            return
        }

        handler(w, r)
    }
}

func main() {
    loginHandler, err := authentication.NewLoginHandler()
    if err != nil {
        log.Fatal("Error loading login handler:", err)
    }

    checkAuthHandler, err := authentication.NewCheckAuthHandler()
    if err != nil {
        log.Fatal("Error loading check auth handler:", err)
    }

    // Public endpoints
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/check_auth", checkAuthHandler)

    // Protected endpoints
    http.HandleFunc("/contractors", withAuth(handlers.AddContractorHandler))
    http.HandleFunc("/lands", withAuth(handlers.AddLandHandler))
    http.HandleFunc("/contracts", withAuth(handlers.AddContractHandler))

    // Public read-only endpoints
    http.HandleFunc("/list_contractors", handlers.ListContractorsHandler)
    http.HandleFunc("/list_lands", handlers.ListLandsHandler)
    http.HandleFunc("/list_contracts", handlers.ListContractsHandler)

    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
