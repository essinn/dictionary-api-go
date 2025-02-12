package handlers

import (
	"dictionary-api/config"
	"dictionary-api/internal/db"
	"dictionary-api/internal/models"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetWord(w http.ResponseWriter, r *http.Request) {
	cfg := config.Load()
	collection := db.GetCollection(cfg)

	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "Missing word parameter", http.StatusBadRequest)
		return 
	}

	var result models.Word
	err := collection.FindOne(r.Context(), bson.M{"word": word}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Word not found", http.StatusNotFound)
		}
		http.Error(w, "Failed to fetch word", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}