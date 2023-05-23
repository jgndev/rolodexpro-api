package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jgndev/rolodexpro-api/internal/dto"
	"github.com/jgndev/rolodexpro-api/internal/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func RegistrationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.RegisterDto

		if err := c.ShouldBindJSON(&request); err != nil {
			log.Printf("Error: binding request failed. %v\n", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Error: error generating hashed password. %v\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user := model.User{
			Email:       request.Email,
			Password:    string(hashedPassword),
			DisplayName: request.Username,
		}

		if err = db.Create(&user).Error; err != nil {
			log.Printf("Error: error creating user in the database. %v\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
	}
}
