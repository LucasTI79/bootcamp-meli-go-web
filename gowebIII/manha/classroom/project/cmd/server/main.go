package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/batatinha123/products-api/cmd/server/handler"
	"github.com/batatinha123/products-api/internal/products"
	"github.com/gin-gonic/gin"
)

// var dbConn *sql.DB

func LoggerMiddleware(ctx *gin.Context) {
	fmt.Printf("[%s] %s\n", ctx.Request.Method, ctx.Request.URL)
	ctx.Next()
}

func TokenMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "123456" {
		// status StatusUnauthorized equivalente ao 401
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token inválido",
		})
		return
	}
	ctx.Next()
}

// middlewares globais
// middlewares de rota

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	productHandler := handler.NewProduct(service)

	server := gin.Default()

	// vai usar o middleware antes de cada handler
	server.Use(LoggerMiddleware)

	pr := server.Group("/products")
	pr.Use(TokenMiddleware)
	pr.POST("/", productHandler.Store())
	pr.GET("/", productHandler.GetAll())
	pr.PUT("/:productId", productHandler.Update())
	pr.PATCH("/:productId", productHandler.UpdateName())
	pr.DELETE("/:productId", productHandler.Delete())

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
