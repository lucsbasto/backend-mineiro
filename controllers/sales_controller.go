package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucsbasto/backend-mineiro/controllers/dtos"
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/services"
)

type SalesController struct {
	service *services.SalesService
}

func NewSalesController(service *services.SalesService) *SalesController {
	return &SalesController{service: service}
}


func (c *SalesController) Create(ctx *gin.Context) {
	var dto dtos.CreateSaleDTO

	// Bind JSON para o DTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tenta pegar o usuário do contexto
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}
	u := user.(models.User) // Agora temos o usuário no formato correto

	// Chama o serviço para criar a venda com o DTO e o userID
	if err := c.service.CreateSale(&dto, u.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna uma resposta de sucesso
	ctx.JSON(http.StatusOK, gin.H{"message": "Sale created successfully"})
}



func (c *SalesController) ListAll(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(401, gin.H{"error": "Usuário não autenticado"})
		return
	}
	u := user.(models.User)
	sales, err := c.service.FindAll(u.IsAdmin, u.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao listar vendas"})
		return
	}
	ctx.JSON(200, sales)
}


func (c *SalesController) ListOne(ctx *gin.Context) {
	
	id := ctx.Param("id")
	sale, err := c.service.FindSaleByID(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Venda não encontrada"})
		return
	}
	ctx.JSON(200, sale)
}

