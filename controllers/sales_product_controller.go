package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/services"
)

type SalesProductController struct {
	service *services.SalesProductService
}

func NewSalesProductController(service *services.SalesProductService) *SalesProductController {
	return &SalesProductController{service: service}
}

func (c *SalesProductController) Create(ctx *gin.Context) {
	var salesProduct models.SalesProduct
	if err := ctx.BindJSON(&salesProduct); err != nil {
		ctx.JSON(400, gin.H{"error": "Erro ao ler dados da linha de venda"})
		return
	}
	if err := c.service.CreateSalesProduct(&salesProduct); err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao criar linha de venda"})
		return
	}
	ctx.JSON(201, gin.H{"message": "Linha de venda criada com sucesso"})
}
