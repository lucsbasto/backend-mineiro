package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucsbasto/backend-mineiro/controllers"
)

// SetupRoutes - Configura as rotas
func SetupRoutes(r *gin.Engine, authController *controllers.AuthController) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", authController.SignIn)
		authGroup.POST("/register", authController.SignUp)
	}
}
