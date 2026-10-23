package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kishanravi887321/pumkin/server-side/internal/handlers"
)

func UserRouter(userRouter *gin.Engine) {
	router := userRouter.Group("/api/user")

	router.POST("/register", handlers.RegisterUser)
	router.GET("/me", handlers.GetMe)
}
