package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const ServerName string = "AppServer"
const ServerPort int = 3000

type AppConfig struct {
	ServerName string
	Port       int
}

func NewAppConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
	}
	log.Println(".env file loaded")
	return &AppConfig{
		ServerName: getString("SERVER_NAME", ServerName),
		Port:       getInt("SERVER_PORT", ServerPort),
	}
}

func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Println("wrong value for server port")
		return defaultValue
	}
	return i
}
