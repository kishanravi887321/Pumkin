package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var user UserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("Username:", user.Username)
	log.Println("Email:", user.Email)

	// Never print passwords in logs.

	c.JSON(http.StatusOK, gin.H{
		"message": "User data received successfully",
	})
}

func GetMe(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "hello Ravi kishan....",
	})
}
