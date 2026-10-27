package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/dto"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
	"github.com/kishanravi887321/pumkin/server-side/internal/services"
	"github.com/kishanravi887321/pumkin/server-side/puzru"
	"github.com/kishanravi887321/pumkin/server-side/requests"
	"github.com/kishanravi887321/pumkin/server-side/utils"
	"google.golang.org/api/idtoken"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user requests.UserRegisterRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	utils.CleanInput(&user)

	createdUser, err := h.userService.Register(c.Request.Context(), dto.RegisterUserRequest{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userClaims := dto.TokenClaims{
		UserID:   createdUser.UserId,
		Username: createdUser.UserName,
		Email:    createdUser.Email,
	}

	accessToken, err := utils.GenerateAccessToken(userClaims)
	// fmt.Println("Access Token:", accessToken) // For demonstration; remove in production.

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate access token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "user registered successfully",
		"user":        createdUser,
		"accessToken": accessToken,
	})
}

func LoginByGoogle(c *gin.Context) {
	var loginRequest requests.UserLoginByGoogleRequest

	cfg, err := config.Load()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Authentication configuration error",
		})
		return
	}
	googleClientID := cfg.GoogleClientID

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fmt.Println("Received Google Token:", loginRequest.GoogleToken)

	payload, err := idtoken.Validate(
		c.Request.Context(),
		loginRequest.GoogleToken,
		googleClientID,
	)
	data := payload.Claims
	fmt.Println("Google Token Payload:", data["email"], data["name"], data["picture"])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Google ID token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Google login successful",
		"payload": payload,
	})

}

func GetMe(c *gin.Context) {

	////
	cfg, err := config.Load()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load configuration",
		})
		return
	}

	puzru.MakeKey()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve encryption key",
		})
		return
	}
	key := cfg.MasterEncryptionKey

	encryptedText, err := puzru.Encrypt("sdvnlkasdnv", []byte(key))

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
		"message":                 "hello Ravi kishan....",
		"decryptedText":           decryptedText,
		"secretFromPrivateServer": response,
	})
}
