package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileHandler(c *gin.Context) {

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user id not found",
		})
		return
	}

	username, _ := c.Get("username")
	clientID, _ := c.Get("client_id")

	c.JSON(http.StatusOK, gin.H{
		"user_id":   userID,
		"username":  username,
		"client_id": clientID,
	})
}
