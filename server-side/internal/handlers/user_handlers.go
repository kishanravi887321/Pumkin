package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/dto"
	"github.com/kishanravi887321/pumkin/server-side/requests"
	"github.com/kishanravi887321/pumkin/server-side/utils"
)

func RegisterUser(c *gin.Context) {
	var user requests.UserRegisterRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("Username:", user.Username)
	log.Println("Email:", user.Email)

	// Never print passwords in logs.

	userClaims := dto.TokenClaims{
		UserID:   "12345", // Replace with the actual ID from your database.
		Username: user.Username,
		Email:    user.Email,
	}

	accessToken, err := utils.GenerateAccessToken(userClaims)
	fmt.Println("Access Token:", accessToken) // For demonstration; remove in production.

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User data received successfully",
	})
}

func GetMe(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "hello Ravi kishan....",
	})
}
