package main

import (
	"fmt"
	"github.com/jgndev/rolodexpro-api/internal/authenticator"
	"github.com/jgndev/rolodexpro-api/internal/config"
	"github.com/jgndev/rolodexpro-api/internal/model"
	"github.com/jgndev/rolodexpro-api/internal/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Starting up...")

	//====================================================================
	// Environment Variables
	//====================================================================
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, will fall back to OS environment variables")
	}

	// Show debug info for environment variables if in debug mode
	if config.Debug {
		config.PrintConfigStatus()
	}

	// Ensure required variables are present before continuing
	for _, v := range config.RequiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Error: required environment variable %v not set\n", v)
		}
	}

	//====================================================================
	// GORM
	//====================================================================
	log.Println("Opening database connection...")
	db, err := gorm.Open("postgres", getDbConnectionStr())
	if err != nil {
		log.Fatalf("Error: failed to connect to database. %v\n", err.Error())
	}
	defer func(db *gorm.DB) {
		err = db.Close()
		if err != nil {
			log.Fatalf("Error: caught closing databsae connection. %v\n", err.Error())
		}
	}(db)

	// Apply migrations
	log.Println("Applying pending migrations...")
	db.AutoMigrate(&model.User{}, &model.Contact{}, &model.Category{})

	//==================================================================
	// Server
	//==================================================================
	log.Println("Configuring authenticator...")
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Error: failed to initialize the authenticator: %v\n", err.Error())
	}

	log.Println("Configuring router...")
	rtr := router.New(auth, db)
	//port := os.Getenv("APP_PORT")
	port := os.Getenv(config.AppPort)
	if err = http.ListenAndServe(":"+port, rtr); err != nil {
		log.Fatalf("Error: failed to start the server. %v\n", err.Error())
	}
}

func getDbConnectionStr() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv(config.DbHost),
		os.Getenv(config.DbPort),
		os.Getenv(config.DbUser),
		os.Getenv(config.DbName),
		os.Getenv(config.DbPassword),
		os.Getenv(config.DbSslMode),
	)
}
