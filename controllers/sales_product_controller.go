package controllers

import (
	"time"

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

func (c *SalesProductController) ListAll(ctx *gin.Context) {
	salesProducts, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao listar vendas"})
		return
	}
	ctx.JSON(200, salesProducts)
}


func (c *SalesProductController) ListByFormattedDate(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(401, gin.H{"error": "Usuário não autenticado"})
		return
	}
	u := user.(models.User)
	
	dateStr := ctx.Param("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Erro ao parsear data"})
		return
	}
	sales, err := c.service.FindSalesByFormattedDate(date.Format("2006-01-02"), u.IsAdmin, u.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao listar vendas"})
		return
	}
	ctx.JSON(200, sales)
}