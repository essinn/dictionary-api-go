package main

import (
	"context"
	"dictionary-api/config"
	"dictionary-api/internal/db"
	"dictionary-api/internal/handlers"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	if err := db.ConnectToMongoDB(cfg); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.Client.Disconnect(context.Background())

	http.HandleFunc("/api/word", handlers.GetWord)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}