package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jgndev/rolodexpro-api/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// Read variables set in local .env file if present
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}

	fmt.Printf("DB_HOST: %v\n", os.Getenv("DB_HOST"))
	fmt.Printf("DB_PORT: %v\n", os.Getenv("DB_PORT"))
	fmt.Printf("DB_USER: %v\n", os.Getenv("DB_USER"))
	fmt.Printf("DB_PASSWORD: %v\n", os.Getenv("DB_PASSWORD"))
	fmt.Printf("DB_NAME: %v\n", os.Getenv("DB_NAME"))
	fmt.Printf("DB_SSLMODE: %v\n", os.Getenv("DB_SSLMODE"))

	// Open the DB connection
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_SSLMODE"),
		),
	)
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	// Apply migrations
	db.AutoMigrate(&model.User{}, &model.Contact{}, &model.Category{})

	// Gin
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
