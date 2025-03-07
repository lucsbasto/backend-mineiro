package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lucsbasto/backend-mineiro/models"
	"gorm.io/gorm"
)


type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (c Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token de autenticação inválido",
			})
			c.Abort()
			return
		}
		
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        
		secretKey := []byte(os.Getenv("JWT_SECRET"))
        
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
        
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token de autenticação inválido",
			})
			c.Abort()
			return
		}
        
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token de autenticação inválido",
			})
			c.Abort()
			return
		}
        
		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token de autenticação inválido",
			})
			c.Abort()
			return
		}
        
		var user models.User
		if err := db.Where(&models.User{Username: claims.Username}).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Usuário não encontrado",
			})
			c.Abort()
			return
		}
        
		c.Set("user", user)
        
		c.Next()
	}
}
