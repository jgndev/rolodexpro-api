package callback

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jgndev/rolodexpro-api/internal/authenticator"
	"log"
	"net/http"
)

func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if ctx.Query("state") != session.Get("state") {
			message := fmt.Sprintf("Error: invalid state parameter.")
			log.Println(message)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": message})
			return
		}

		// Exchange an authorization code for a token.
		token, err := auth.Exchange(ctx.Request.Context(), ctx.Query("code"))
		if err != nil {
			message := fmt.Sprintf("Error: failed to exchange an authorization code for a token. %v\n", err.Error())
			log.Println(message)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": message})
			return
		}

		idToken, err := auth.VerifyIDToken(ctx.Request.Context(), token)
		if err != nil {
			message := fmt.Sprintf("Error: failed to verify ID token. %v\n", err.Error())
			log.Println(message)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
			return
		}

		var profile map[string]interface{}
		if err = idToken.Claims(&profile); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		session.Set("access_token", token.AccessToken)
		session.Set("profile", profile)
		if err = session.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		// Redirect to the home page
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
