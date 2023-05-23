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

func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.LoginDto

		if err := c.ShouldBindJSON(&request); err != nil {
			log.Printf("Error: binding request failed. %v\n", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user model.User
		if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
			log.Printf("Error: provided email does not exist. %v\n", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password"})
			return
		}

		// TODO: Generate a token and return it in the response
		c.JSON(http.StatusOK, gin.H{"message": "User authenticated"})
	}
}
