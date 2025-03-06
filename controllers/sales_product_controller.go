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

func (c *SalesController) CreateSalesProduct(ctx *gin.Context) {
	var sale models.Sales
	if err := ctx.BindJSON(&sale); err != nil {
		ctx.JSON(400, gin.H{"error": "Erro ao ler dados da venda"})
		return
	}

	// Recupera o usuário autenticado do contexto
	user, _ := ctx.Get("user")
	u := user.(models.User)

	// Associa o usuário à venda
	sale.UserID = u.ID

	// Salva os produtos da venda
	for i := range sale.SalesProducts {
		sale.SalesProducts[i].SaleID = sale.ID
	}

	if err := c.service.CreateSale(&sale); err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao criar venda"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Venda criada com sucesso"})
}