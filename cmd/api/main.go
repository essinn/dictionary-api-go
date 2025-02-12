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
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}