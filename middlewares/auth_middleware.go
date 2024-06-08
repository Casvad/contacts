package middlewares

import (
	"contacts/dto"
	"contacts/utils/constants"
	"contacts/utils/env"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		if tokenString == "" {
			c.IndentedJSON(http.StatusUnauthorized, map[string]string{
				"error": "Empty Authorization Header",
			})
			c.Abort()
			return
		}

		claims := &dto.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.JwtKey), nil
		})

		if err != nil || !token.Valid {
			c.IndentedJSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set(constants.Claims, claims)
		c.Next()
	}
}
