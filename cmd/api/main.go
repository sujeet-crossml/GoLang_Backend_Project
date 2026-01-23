package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sujeet-crossml/GoLang_Backend_Project/internal/config"
	"github.com/sujeet-crossml/GoLang_Backend_Project/internal/handlers"
	"github.com/sujeet-crossml/GoLang_Backend_Project/internal/middleware"
)

func main() {
	// Initialize DB
	config.ConnectDB()

	//Router setup
	mux := http.NewServeMux()

	// Public Router
	mux.HandleFunc("POST /register", handlers.Register)
	mux.HandleFunc("POST /login", handlers.Login)

	// Protected routes (authmiddleware)
	mux.HandleFunc("GET /profile", middleware.AuthMiddleware(handlers.GetProfile))
	mux.HandleFunc("GET /orders", middleware.AuthMiddleware(handlers.GetMyOrders))

	// Start server
	port := ":8080"

	fmt.Printf("Server running on port %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
