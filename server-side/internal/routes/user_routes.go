package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/internal/handlers"
	"github.com/kishanravi887321/pumkin/server-side/internal/middleware"
)

func UserRouter(userRouter *gin.Engine) {

	//  these are public routes, no need to validate token
	publicRoute := userRouter.Group("/api/user")
	publicRoute.POST("/register", handlers.RegisterUser)

	protectedRoute := userRouter.Group("/api/user")
	protectedRoute.Use(middleware.AuthMiddleware())
	protectedRoute.GET("/me", handlers.GetMe)

	// Add protected routes here, e.g.:
	// protectedRoute.GET("/profile", handlers.GetUserProfile)
}
