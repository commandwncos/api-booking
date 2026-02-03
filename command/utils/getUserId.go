package utils

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) (int64, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}

	id, ok := userID.(int64)
	return id, ok
}
