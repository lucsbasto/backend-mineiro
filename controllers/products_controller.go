package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/services"
)

// ProductController encapsula a lógica de controle para produtos.
type ProductController struct {
	service *services.ProductService
}

// NewProductController cria uma nova instância de ProductController.
func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// ListAll lista todos os produtos.
func (c *ProductController) ListAll(ctx *gin.Context) {
	products, err := c.service.FindAllProducts()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao listar produtos"})
		return
	}
	ctx.JSON(200, products)
}

// ListOne busca um produto por ID.
func (c *ProductController) ListOne(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := c.service.FindProductByID(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Produto não encontrado"})
		return
	}
	ctx.JSON(200, product)
}

// Create insere um novo produto.
func (c *ProductController) Create(ctx *gin.Context) {
	var product models.Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": "Erro ao ler dados do produto"})
		return
	}
	if err := c.service.CreateProduct(&product); err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao criar produto"})
		return
	}
	ctx.JSON(201, gin.H{"message": "Produto criado com sucesso"})
}
// Update atualiza um produto existente.
func (c *ProductController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": "Erro ao ler dados do produto"})
		return
	}

	// Busca o produto existente
	existingProduct, err := c.service.FindProductByID(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Produto não encontrado"})
		return
	}

	// Atualiza o ID do produto recebido para garantir que seja o mesmo produto
	product.ID = existingProduct.ID

	if err := c.service.UpdateProduct(&product); err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao atualizar produto"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Produto atualizado com sucesso"})
}


// Delete exclui um produto logicamente.
func (c *ProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteProduct(id); err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao excluir produto"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Produto excluído com sucesso"})
}
