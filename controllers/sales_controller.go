package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
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
	var sale models.Sales
	if err := ctx.BindJSON(&sale); err != nil {
		ctx.JSON(400, gin.H{"error": "Erro ao ler dados da venda"})
		return
	}

	// Recupera o usuário autenticado do contexto
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(401, gin.H{"error": "Usuário não autenticado"})
		return
	}
	u := user.(models.User)

	// Associa o usuário à venda
	sale.UserID = u.ID

	if err := c.service.CreateSaleWithProducts(&sale); err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao criar venda"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Venda criada com sucesso"})
}

func (c *SalesController) ListAll(ctx *gin.Context) {
	sales, err := c.service.FindAll()
	fmt.Println(sales)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao listar vendas"})
		return
	}
	ctx.JSON(200, sales)
}

func (c *SalesController) ListByFormattedDate(ctx *gin.Context) {
	dateStr := ctx.Param("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Erro ao parsear data"})
		return
	}
	sales, err := c.service.FindSalesByFormattedDate(date.Format("2006-01-02"))
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

func (c *SalesController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var sale models.Sales
	if err := ctx.BindJSON(&sale); err != nil {
		ctx.JSON(400, gin.H{"error": "Erro ao ler dados da venda"})
		return
	}

	existingSale, err := c.service.FindSaleByID(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Venda não encontrada"})
		return
	}

	sale.ID = existingSale.ID
	if err := c.service.UpdateSale(&sale); err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao atualizar venda"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Venda atualizada com sucesso"})
}

func (c *SalesController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteSale(id); err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao excluir venda"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Venda excluída com sucesso"})
}
