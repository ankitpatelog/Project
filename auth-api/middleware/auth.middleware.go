package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get the header
		auth := c.GetHeader("Authorization")

		//chechk if header conans the bearer word it means it has valid syntax
		if !strings.HasPrefix(auth,"Bearer") {
			c.AbortWithStatusJSON(401,gin.H{"error": "Unauthorized"})
			return
		}

		tokenstr := strings.TrimPrefix(auth,"Bearer")

		token, err := jwt.Parse(tokenstr, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err!=nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		c.Next()
	}
}


