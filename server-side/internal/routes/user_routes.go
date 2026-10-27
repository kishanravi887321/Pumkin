package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/internal/handlers"
	"github.com/kishanravi887321/pumkin/server-side/internal/middleware"
)

func UserRouter(userRouter *gin.Engine, userHandler *handlers.UserHandler) {

	//  these are public routes, no need to validate token
	publicRoute := userRouter.Group("/api/user")
	publicRoute.POST("/register", userHandler.RegisterUser)
	publicRoute.POST("/google-login", handlers.LoginByGoogle)

	protectedRoute := userRouter.Group("/api/user")
	protectedRoute.Use(middleware.AuthMiddleware())
	protectedRoute.GET("/me", handlers.GetMe)

	// Add protected routes here, e.g.:
	// protectedRoute.GET("/profile", handlers.GetUserProfile)
}
