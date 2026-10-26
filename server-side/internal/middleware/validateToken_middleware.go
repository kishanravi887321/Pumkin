package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kishanravi887321/pumkin/server-side/internal/config"
)

// AuthMiddleware validates the access token before allowing
// the request to reach a protected handler.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Step 1: Read the Authorization header.
		// Expected format: Authorization: Bearer <access_token>
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			return
		}

		// Step 2: Split the header into the scheme and token.
		parts := strings.Fields(authHeader)

		// The header must contain exactly two parts:
		// "Bearer" and the JWT.
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format",
			})
			return
		}

		tokenString := parts[1]

		// Step 3: Get the secret used to sign access tokens.
		secret, err := config.GetAccessTokenSecret()
		if err != nil {
			// A configuration error is a server error,
			// not an invalid-token error.
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Authentication configuration error",
			})
			return
		}

		// Step 4: Parse and validate the JWT.
		// jwt.Parse checks the signature and registered claims,
		// including expiration when an exp claim is present.
		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (any, error) {

				// Accept only the expected HS256 signing method.
				if token.Method != jwt.SigningMethodHS256 {
					return nil, jwt.ErrSignatureInvalid
				}

				// Return the secret used to verify the signature.
				return []byte(secret), nil
			},
			jwt.WithValidMethods([]string{"HS256"}),
		)

		// Reject malformed tokens, invalid signatures,
		// expired tokens, and other validation failures.
		if err != nil || token == nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired access token",
			})
			return
		}

		// Step 5: Extract the claims from the validated token.
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			return
		}

		// Ensure the required user ID claim exists.
		userID, ok := claims["user_id"].(string)
		if !ok || userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid user identity in token",
			})
			return
		}

		// Step 6: Store validated claims in Gin's context.
		// Protected handlers can retrieve these values using c.Get().
		c.Set("userId", userID)

		if username, ok := claims["username"].(string); ok {
			c.Set("userName", username)
		}

		if email, ok := claims["email"].(string); ok {
			c.Set("email", email)
		}

		// Step 7: Continue to the next middleware or handler.
		c.Next()
	}
}
