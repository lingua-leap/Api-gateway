package middleware

import (
	"api-gateway/pkg/token"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		stringToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, err := token.ExtractClaims(stringToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func AuthorizationMiddleware(csb *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		stringToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
		}

		claims, err := token.ExtractClaims(stringToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
		}

		check, err := csb.Enforce(claims.Role, c.FullPath(), c.Request.Method)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
		}

		if !check {
			c.JSON(http.StatusUnauthorized, gin.H{"UnAuthorization": "access denied"})
			c.Abort()
		}

		c.Next()
	}
}
