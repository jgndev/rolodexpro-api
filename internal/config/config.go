package config

import (
	"log"
	"os"
)

const (
	DbHost            = "DB_HOST"
	DbPort            = "DB_PORT"
	DbUser            = "DB_USER"
	DbPassword        = "DB_PASSWORD"
	DbName            = "DB_NAME"
	DbSslMode         = "DB_SSLMODE"
	AppPort           = "APP_PORT"
	Auth0Domain       = "AUTH0_DOMAIN"
	Auth0ClientId     = "AUTH0_CLIENT_ID"
	Auth0ClientSecret = "AUTH0_CLIENT_SECRET"
	Auth0CCallbackUrl = "AUTH0_CALLBACK_URL"
	JwtSecret         = "JWT_SECRET"
	Debug             = true
)

var RequiredVars = []string{
	DbHost,
	DbPort,
	DbUser,
	DbPassword,
	DbName,
	DbSslMode,
	AppPort,
	Auth0Domain,
	Auth0ClientId,
	Auth0ClientSecret,
	Auth0CCallbackUrl,
	JwtSecret,
}

func PrintConfigStatus() {
	log.Println("Checking state of required environment variables")
	for _, v := range RequiredVars {
		if os.Getenv(v) != "" {
			log.Printf("%-20s: %t\n", v, true)
		} else {
			log.Printf("%-20s: %t\n", v, false)
		}
	}
}
