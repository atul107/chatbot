package middleware

import (
	"net/http"
	"strings"

	"github.com/chatbot/internal/model"
	"github.com/gin-gonic/gin"
)

// ParseTokenFromHeaderOrCookie ...
func ParseTokenFromHeaderOrCookie(c *gin.Context) (string, bool) {
	// Check if the Authorization header has the Bearer prefix
	BearerToken := c.Request.Header.Get("Authorization")
	if BearerToken != "" {
		parts := strings.SplitN(BearerToken, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1], true
		}
	}

	return "", false
}

// Authenticate ...
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := ParseTokenFromHeaderOrCookie(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Message: "Unauthorized"})
			return
		}
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Message: "Unauthorized"})
			return
		}
		c.Next()
	}
}
