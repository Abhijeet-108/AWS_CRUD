package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CognitoRegion       string
	CognitoUserPoolID   string
	CognitoDomain       string
	CognitoClientID     string
	CognitoClientSecret string
	CognitoRedirectURI  string
	EmployeeAPIURL      string
}

func Load() Config {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return Config{
		CognitoRegion:       getConfig("COGNITO_REGION", "us-east-1"),
		CognitoUserPoolID:   getConfig("COGNITO_USER_POOL_ID", ""),
		CognitoDomain:       getConfig("COGNITO_DOMAIN", ""),
		CognitoClientID:     getConfig("COGNITO_CLIENT_ID", ""),
		CognitoClientSecret: getConfig("COGNITO_CLIENT_SECRET", ""),
		CognitoRedirectURI:  getConfig("COGNITO_REDIRECT_URI", ""),
		EmployeeAPIURL:      getConfig("EMPLOYEE_API_URL", ""),
	}

}

func getConfig(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
