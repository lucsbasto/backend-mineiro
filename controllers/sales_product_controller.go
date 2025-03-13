package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucsbasto/backend-mineiro/helpers"
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/services"
	"github.com/lucsbasto/backend-mineiro/types"
)

type SalesProductController struct {
	service *services.SalesProductService
}

func NewSalesProductController(service *services.SalesProductService) *SalesProductController {
	return &SalesProductController{service: service}
}

func (c *SalesProductController) UpdateSaleProduct(ctx *gin.Context) {
	
}

func (c *SalesProductController) ListAll(ctx *gin.Context) {
	salesProducts, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao listar vendas"})
		return
	}
	ctx.JSON(200, salesProducts)
}

func (c *SalesProductController) Update(ctx *gin.Context) {
	user, err := c.getUserFromContext(ctx)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("id")
	var dto types.UpdateSalesProduct

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(400, gin.H{"error": "Erro ao fazer bind do JSON"})
		return
	}

	salesProduct, err := c.service.ListOne(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao buscar venda"})
		return
	}
	fmt.Println("salesProduct: ", salesProduct)

	if user.IsAdmin {
		salesProduct.Quantity = helpers.Coalesce(*dto.Quantity, salesProduct.Quantity)
		salesProduct.Sold = helpers.Coalesce(*dto.Sold, salesProduct.Sold)
		salesProduct.Returned = helpers.Coalesce(*dto.Returned, salesProduct.Returned)
		salesProduct.UnitCost = helpers.CoalesceFloat64(*dto.UnitCost, salesProduct.UnitCost)
		salesProduct.Price = helpers.CoalesceFloat64(*dto.Price, salesProduct.Price)
	} else {
		salesProduct.Returned = helpers.Coalesce(*dto.Returned, salesProduct.Returned)
		salesProduct.Sold = helpers.Coalesce(*dto.Sold, salesProduct.Sold)
	}
	if err := c.service.Update(salesProduct); err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao atualizar venda"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Venda atualizada com sucesso"})
}


func (c *SalesProductController) ListByFormattedDate(ctx *gin.Context) {
	user, err := c.getUserFromContext(ctx)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	dateStr := ctx.Param("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Erro ao parsear data"})
		return
	}
	sales, err := c.service.FindSalesByFormattedDate(date.Format("2006-01-02"), user.IsAdmin, user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao listar vendas"})
		return
	}
	ctx.JSON(200, sales)
}

func (c *SalesProductController) ListOne(ctx *gin.Context) {
	id := ctx.Param("id")
	salesProduct, err := c.service.ListOne(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Venda não encontrada"})
		return
	}
	ctx.JSON(200, salesProduct)
}

func (c *SalesProductController) getUserFromContext(ctx *gin.Context) (*models.User, error) {
	user, ok := ctx.Get("user")
	if !ok {
		return nil, errors.New("usuário não autenticado")
	}
	u, ok := user.(models.User)
	if !ok {
		return nil, errors.New("usuário não é do tipo models.User")
	}
	return &u, nil
}
