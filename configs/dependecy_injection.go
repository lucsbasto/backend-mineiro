package configs

import (
	"github.com/lucsbasto/backend-mineiro/controllers"
	"github.com/lucsbasto/backend-mineiro/database"
	"github.com/lucsbasto/backend-mineiro/repositories"
	"github.com/lucsbasto/backend-mineiro/services"
)

func InitializeDependencies() (*controllers.AuthController, error) {
	Init()
	
	database.Connect()

	userRepo := repositories.NewUserRepository(database.DB)

	authService := services.NewAuthService(userRepo)

	authController := controllers.NewAuthController(authService)

	return authController, nil
}
