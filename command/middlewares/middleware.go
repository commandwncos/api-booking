package middlewares

import (
	"log"
	"strings"

	"github.com/commandwncos/api-booking/command/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	log.Println("AUTH MIDDLEWARE HIT")
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "missing authorization header"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"message": "invalid authorization format"})
			return
		}

		claims, err := utils.VerifyJsonWebToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "invalid token"})
			return
		}

		// 🔥 ISSO É O QUE SEU HANDLER ESPERA
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
