package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jgndev/rolodexpro-api/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var (
	DEBUG         = true
	REQUIRED_VARS = []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSLMODE", "APP_PORT"}
)

func main() {

	startTime := time.Now()
	log.Printf("Starting up at %v\n", startTime)

	//====================================================================
	// Environment Variables
	//====================================================================
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, will fall back to OS environment variables")
	}

	// Show debug info for environment variables if in debug mode
	if DEBUG {
		printEnvDebugInfo()
	}

	// Ensure required variables are present before continuing
	for _, v := range REQUIRED_VARS {
		if os.Getenv(v) == "" {
			log.Fatalf("Error: required environment variable %v not set\n", v)
		}
	}

	//====================================================================
	// GORM
	//====================================================================
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
	db.AutoMigrate(&model.User{}, &model.Contact{}, &model.Category{})

	//==================================================================
	// Gin
	//==================================================================
	r := gin.Default()
	if DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	port := os.Getenv("APP_PORT")
	// Fall back to port 8080 in the event we made it here and APP_PORT is not set
	if port == "" {
		port = "8080"
	}
	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error: failed to start the server. %v\n", err.Error())
	}
}

func printEnvDebugInfo() {
	for _, v := range REQUIRED_VARS {
		if os.Getenv(v) != "" {
			log.Printf("%-16s %v\n", v+" set:", true)
		} else {
			log.Printf("%-16s %v\n", v+" set:", false)
		}
	}
}

func getDbConnectionStr() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSLMODE"),
	)
}
