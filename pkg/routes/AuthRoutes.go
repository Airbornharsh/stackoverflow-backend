package routes

import (
	"github.com/airbornharsh/stackoverflow-backend/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutesInit(r *gin.RouterGroup) {

	auth := r.Group("/auth")

	auth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Auth Routes",
		})
	})

	auth.POST("/register", handlers.RegisterHandler)
	auth.POST("/login", handlers.LoginHandler)
	// r.POST("/auth/logout", Logout)
	// r.POST("/auth/refresh", Refresh)
	// r.POST("/auth/forgot-password", ForgotPassword)
	// r.POST("/auth/reset-password", ResetPassword)
}
