package main

import (
	"log"

	"github.com/batatinha123/products-api/cmd/server/handler"
	"github.com/batatinha123/products-api/internal/products"
	"github.com/gin-gonic/gin"
)

// var dbConn *sql.DB

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	productHandler := handler.NewProduct(service)

	server := gin.Default()
	pr := server.Group("/products")
	pr.POST("/", productHandler.Store())
	pr.GET("/", productHandler.GetAll())
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
