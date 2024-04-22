package main

import (
	"log"

	"github.com/batatinha123/bootcamp-meli-web/cmd/server/handler"
	"github.com/batatinha123/bootcamp-meli-web/cmd/server/middleware"
	"github.com/batatinha123/bootcamp-meli-web/internal/products"
	"github.com/batatinha123/bootcamp-meli-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	db := store.NewFileStore("file", "products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	productHandler := handler.NewProduct(service)

	server := gin.Default()

	pr := server.Group("/products")
	pr.Use(middleware.TokenAuthMiddleware())
	pr.POST("/", productHandler.Store())
	pr.GET("/", productHandler.GetAll())
	pr.PUT("/:productId", productHandler.Update())
	pr.PATCH("/:productId", productHandler.UpdateName())
	pr.DELETE("/:productId", productHandler.Delete())

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
