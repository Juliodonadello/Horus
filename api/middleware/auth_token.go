package middleware

import (
	"api/helpers"
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenDeviceAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		bearerToken := c.Request.Header.Get("Authorization")
		if helpers.ValidateBearerToken(bearerToken) {
			token = helpers.ExtractBerarerToken(bearerToken)
		} else {
			c.Header("WWW-Authenticate", "Malformed token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		isValid, _ := models.CheckToken(token)
		if !isValid {
			c.Header("WWW-Authenticate", "Authorization Required")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
