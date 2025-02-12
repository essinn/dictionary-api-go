package config

import (
	"os"
)

type Config struct {
	Port string
	MongoDBURI string
	DBName string
	Collection string
}

func Load() *Config {
	return &Config{
		Port: os.Getenv("PORT"),
		MongoDBURI: os.Getenv("MONGODB_URI"),
		DBName: os.Getenv("DB_NAME"),
		Collection: os.Getenv("COLLECTION"),
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

