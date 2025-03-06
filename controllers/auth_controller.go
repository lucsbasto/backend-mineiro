package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucsbasto/backend-mineiro/services"
	"github.com/lucsbasto/backend-mineiro/types"
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
	var signInRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name string `json:"name"`
	}
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

func (ctrl *AuthController) SignUp(ctx *gin.Context){
	var data types.SignUpDTO
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ctrl.authService.SignUp(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": user.ID, "username": user.Username})
}