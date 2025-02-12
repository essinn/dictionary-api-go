package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	MongoDBURI string
	DBName string
	Collection string
}

func Load() *Config {
    cfg := &Config{
        Port:       GetEnv("PORT", "3000"),
        MongoDBURI: GetEnv("MONGODB_URI", ""),
        DBName:     GetEnv("DB_NAME", "dictionary"),
        Collection: GetEnv("COLLECTION", "words"),
    }
    
    fmt.Printf("Loaded config: %+v\n", cfg)
    return cfg
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

