package router

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jgndev/rolodexpro-api/internal/authenticator"
	"github.com/jgndev/rolodexpro-api/internal/callback"
	"github.com/jgndev/rolodexpro-api/internal/config"
	"github.com/jgndev/rolodexpro-api/internal/handler"
	"github.com/jinzhu/gorm"
	"net/http"
)

func New(auth *authenticator.Authenticator, db *gorm.DB) *gin.Engine {
	router := gin.Default()
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// For registering types in cookies
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("authenticator-session", store))

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.POST("/auth/register", handler.RegistrationHandler(db))
	router.POST("/auth/login", handler.LoginHandler(db))
	router.GET("/auth0/login", handler.AuthZeroLoginHandler(db, auth))
	router.GET("/auth0/logout", handler.AuthZeroLogoutHandler())
	router.GET("/callback", callback.Handler(auth))
	//router.GET("/user", middleware.IsAuthenticated, user.Handler)
	//router.GET("/auth0/login", authZero.CallbackHandler(db))
	//router.GET("/auth0/callback", authZero.CallbackHandler(db))
	//router.GET("/user", middleware.JwtMiddleware(), user.Handler)

	return router
}
