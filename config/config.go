package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SlackWebhookFinanceBugs string
	Port                    string
	Environment             string
	Host                    string
	KeycloakURL             string
	KeycloakClientId        string
	KeycloakClientSecret    string
	KeycloakRealm           string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	return Config{
		SlackWebhookFinanceBugs: getEnv("SLACK_WEBHOOK_CORE_BUGS", ""),
		Port:                    getEnv("PORT", "8080"),
		Environment:             getEnv("ENVIRONMENT", "dev"),
		Host:                    getEnv("HOST", "localhost"),
		KeycloakURL:             getEnv("KEYCLOAK_URL", ""),
		KeycloakClientId:        getEnv("KEYCLOAK_CLIENT_ID", ""),
		KeycloakClientSecret:    getEnv("KEYCLOAK_CLIENT_SECRET", ""),
		KeycloakRealm:           getEnv("KEYCLOAK_REALM", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
