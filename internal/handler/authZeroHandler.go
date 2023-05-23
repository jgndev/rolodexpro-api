package handler

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jgndev/rolodexpro-api/internal/authenticator"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/url"
	"os"
)

func AuthZeroLoginHandler(db *gorm.DB, auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Save the session inside the session
		session := sessions.Default(ctx)
		session.Set("state", state)
		if err = session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))
	}
}

func AuthZeroLogoutHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		scheme := "http"
		if ctx.Request.TLS != nil {
			scheme = "https"
		}

		returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		parameters := url.Values{}
		parameters.Add("returnTo", returnTo.String())
		parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
		logoutUrl.RawQuery = parameters.Encode()

		ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
