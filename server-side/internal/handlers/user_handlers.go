package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/dto"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
	"github.com/kishanravi887321/pumkin/server-side/puzru"
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
	utils.CleanInput(&user)

	fmt.Print(user.Username, user.Email)

	// Never print passwords in logs.

	userClaims := dto.TokenClaims{
		UserID:   "12345", // Replace with the actual ID from your database.
		Username: user.Username,
		Email:    user.Email,
	}

	accessToken, err := utils.GenerateAccessToken(userClaims)
	// fmt.Println("Access Token:", accessToken) // For demonstration; remove in production.

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "User data received successfully",
		"accessToken": accessToken,
	})
}

func GetMe(c *gin.Context) {

	////
	key, err := config.GetEncryptionKey()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve encryption key",
		})
		return
	}

	encryptedText, err := puzru.Encrypt("postgresql://neondb_owner:npg_ZbsXM1uRahJ2@ep-winter-resonance-az3jo11s-pooler.c-3.ap-southeast-1.aws.neon.tech/neondb?sslmode=require", []byte(key))

	fmt.Println("Encrypted Text size :", len(encryptedText))
	if err != nil {
		log.Println("Encryption error:", err)
		return
	}

	fmt.Println("Encrypted:", encryptedText)
	if err != nil {
		fmt.Println("Encryption error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to encrypt text",
		})
		return
	}

	fmt.Println("Encrypted Text:", encryptedText)

	decryptedText, err := puzru.Decrypt(encryptedText, []byte(key))
	if err != nil {
		fmt.Println("Decryption error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decrypt text",
		})
		return
	}

	fmt.Println("Decrypted Text:", decryptedText)

	response, err := puzru.CallPrivateServer()
	fmt.Println("Response from private server:", response)

	if err != nil {
		fmt.Println("Error calling private server:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to call private server",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":       "hello Ravi kishan....",
		"decryptedText": decryptedText,
	})
}
