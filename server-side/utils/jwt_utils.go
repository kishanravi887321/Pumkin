package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kishanravi887321/pumkin/server-side/dto"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
)

func GenerateAccessToken(userClaims dto.TokenClaims) (string, error) {

	secret, err := config.GetAccessTokenSecret()
	if err != nil {
		return "", err
	}

	expiry, err := config.GetAccessTokenExpiry()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userClaims.UserID,
		"username": userClaims.Username,
		"email":    userClaims.Email,
		"exp":      time.Now().Add(time.Duration(expiry)).Unix(),
	})

	return token.SignedString([]byte(secret))
}
