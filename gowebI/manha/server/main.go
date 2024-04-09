package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// inicializando uma instancia do gin com os middlewares default
	router := gin.Default()

	// criando uma rota para o método GET
	router.GET("/hello-world", func(c *gin.Context) {
		router.GET("/hello-world", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World!",
			})
		})

	})

	// rodando o servidor na porta 8080
	router.Run(":8080")
}
