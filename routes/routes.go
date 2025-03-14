package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lucsbasto/backend-mineiro/controllers"
	"github.com/lucsbasto/backend-mineiro/middlewares"
	"gorm.io/gorm"
)

// Controllers é uma estrutura para agrupar todos os controllers.
type Controllers struct {
	AuthController         *controllers.AuthController
	ProductController      *controllers.ProductController
	SalesController        *controllers.SalesController
	SalesProductController *controllers.SalesProductController
}

// SetupRoutes configura as rotas para o servidor.
func SetupRoutes(r *gin.Engine, controllers Controllers, db *gorm.DB) {
	// Configurações CORS para permitir todas as origens
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true   // Permitir todas as origens
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "PATCH"} // Permitir todos os métodos
	config.AllowHeaders = []string{"Content-Type", "Authorization", "X-Requested-With"} // Permitir todos os cabeçalhos
	config.AllowCredentials = true // Permitir cookies (se necessário)

	// Aplicar CORS em todas as rotas
	r.Use(cors.New(config))

	// Rotas para autenticação (públicas)
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", controllers.AuthController.SignIn)
		authRoutes.POST("/register", controllers.AuthController.SignUp)
	}

	// Grupo de rotas protegidas
	protectedRoutes := r.Group("/")
	protectedRoutes.Use(middlewares.AuthMiddleware(db))

	// Rotas para produtos
	productRoutes := protectedRoutes.Group("/products")
	{
		productRoutes.GET("/", controllers.ProductController.ListAll)
		productRoutes.GET("/:id", controllers.ProductController.ListOne)
		productRoutes.POST("/", controllers.ProductController.Create)
		productRoutes.PUT("/:id", controllers.ProductController.Update)
		productRoutes.DELETE("/:id", controllers.ProductController.Delete)
	}

	// Rotas para vendas
	salesRoutes := protectedRoutes.Group("/sales")
	{
		salesRoutes.GET("/", controllers.SalesController.ListAll)
		// salesRoutes.GET("/:id", controllers.SalesController.ListOne)
		salesRoutes.POST("/", controllers.SalesController.Create)
		// salesRoutes.PUT("/:id", controllers.SalesController.Update)
		// salesRoutes.DELETE("/:id", controllers.SalesController.Delete)
	}

	// Rotas para vendas de produtos
	salesProductsRoutes := protectedRoutes.Group("/sales-products")
	{
		salesProductsRoutes.GET("/", controllers.SalesProductController.ListAll)
		salesProductsRoutes.GET("/by-date/:date", controllers.SalesProductController.ListByFormattedDate)
		salesProductsRoutes.GET("/:id", controllers.SalesProductController.ListOne)
		salesProductsRoutes.PATCH("/:id", controllers.SalesProductController.Update)
	}
}
