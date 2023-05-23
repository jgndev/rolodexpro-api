package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAuthenticated(ctx *gin.Context) {
	if sessions.Default(ctx).Get("profile") == nil {
		ctx.Redirect(http.StatusSeeOther, "/")
	} else {
		ctx.Next()
	}
}
