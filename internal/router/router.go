package router

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jgndev/rolodexpro-api/internal/authenticator"
	"github.com/jgndev/rolodexpro-api/internal/callback"
	"github.com/jgndev/rolodexpro-api/internal/handler"
	"github.com/jgndev/rolodexpro-api/internal/middleware"
	"github.com/jgndev/rolodexpro-api/internal/user"
	"github.com/jinzhu/gorm"
	"net/http"
)

func New(auth *authenticator.Authenticator, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// For registering types in cookies
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("authenticator-session", store))

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.POST("/register", handler.RegistrationHandler(db))
	router.POST("/login", handler.LoginHandler(db))
	router.GET("/authenticator-zero-login", handler.AuthZeroLoginHandler(db, auth))
	router.GET("/authenticator-zero-logout", handler.AuthZeroLogoutHandler())
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)

	return router
}
