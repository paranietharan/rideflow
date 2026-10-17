package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	AppEnv      string
	MongoURI    string
	MongoDBName string
	JWTSecret   string
}

func Load() *Config {
	_ = godotenv.Load(".env", "./.env")

	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8000"),
		AppEnv:      getEnv("APP_ENV", "development"),
		MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDBName: getEnv("MONGO_DB_NAME", "rideflow"),
		JWTSecret:   getEnv("JWT_SECRET", "rideflow-dev-secret-key"),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
