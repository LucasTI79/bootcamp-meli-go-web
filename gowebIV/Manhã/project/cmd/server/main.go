package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/batatinha123/products-api/cmd/server/handler"
	"github.com/batatinha123/products-api/internal/products"
	"github.com/batatinha123/products-api/pkg/store"
	"github.com/batatinha123/products-api/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// var dbConn *sql.DB

func LoggerMiddleware(ctx *gin.Context) {
	fmt.Printf("[%s] %s\n", ctx.Request.Method, ctx.Request.URL)
	ctx.Next()
}

func TokenMiddleware(ctx *gin.Context) {
	tokenEnvironment := os.Getenv("TOKEN")
	token := ctx.GetHeader("token")
	if token != tokenEnvironment {
    ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(int(http.StatusBadRequest), nil, "token inválido"))
		return
	}
	ctx.Next()
}

// middlewares globais
// middlewares de rota

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	user := os.Getenv("MY_USER")
	password := os.Getenv("MY_PASS")

	fmt.Println("user", user, "pass", password)

	db := store.NewFileStore("file", "products.json")
	repo := products.NewRepository(db)
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
