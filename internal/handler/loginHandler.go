package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jgndev/rolodexpro-api/internal/dto"
	"github.com/jgndev/rolodexpro-api/internal/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
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

		claims := &jwtCustomClaims{
			User: user,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		ss, err := token.SignedString(jwtSecret)
		if err != nil {
			message := fmt.Sprintf("Error: could not generate token. %v\n", err.Error())
			log.Println(message)
			c.JSON(http.StatusInternalServerError, gin.H{"error": message})
			return
		}

		// TODO: Generate a token and return it in the response
		//c.JSON(http.StatusOK, gin.H{"message": "User authenticated"})
		c.JSON(http.StatusOK, gin.H{"token": ss})
	}
}
