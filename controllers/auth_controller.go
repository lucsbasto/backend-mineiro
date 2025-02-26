package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/services"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ctrl *AuthController) SignIn(c *gin.Context) {
	var signInRequest models.SignInRequest
	err := c.ShouldBindJSON(&signInRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	token, err := ctrl.authService.SignIn(signInRequest.Username, signInRequest.Password)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ctrl *AuthController) SignUp(c *gin.Context){
	var signUpRequest models.SignUpRequest
	err := c.ShouldBindJSON(&signUpRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := ctrl.authService.SignUp(signUpRequest.Username, signUpRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}