package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type produto struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Tipo       string  `json:"tipo"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}

var products []produto
var lastId = 0

// requisição -> processamento -> resposta
// recebemos a request
// guardamos o conteudo do body da request numa variável com o ShouldBindJson()
// aplicamos regras que forem necessárias
// retornamos a response para o cliente com c.Json()

func Salvar() gin.HandlerFunc {
	return func(c *gin.Context) {
		// os query params e route params são case sensitive
		batata := c.Query("batata")
		fmt.Printf("param batata: %s \n", batata)

		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token não informado",
			})
			return
		}

		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token inválido",
			})
			return
		}

		var req produto
		// podemos fazer relação com o json.Unmarshal([]byte(c.Request.Body), &req)
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastId++
		req.ID = lastId
		products = append(products, req)
		c.JSON(http.StatusCreated, req)
	}
}

func Update(c *gin.Context) {
	var req produto

	// podemos fazer relação com o json.Unmarshal([]byte(c.Request.Body), &req)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func main() {
	server := gin.Default()

	// /api/products
	// GET, POST
	// /api/products/:id
	// GET, PUT, PATCH, DELETE

	productsGroup := server.Group("/produtos")

	// se fossemos chamar essa rota, seria /produtos/
	productsGroup.GET("/", func(c *gin.Context) {})
	// se fossemos chamar essa rota, seria /produtos/
	productsGroup.POST("/", Salvar(), func(c *gin.Context) {})
	// se fossemos chamar essa rota, seria /produtos/:productId
	productsGroup.GET("/:productId", func(c *gin.Context) {
		// para capturar esse parâmetro de rota, usamos o c.Param("<nome_parametro>")
		productId := c.Param("productId")
		c.JSON(http.StatusOK, gin.H{
			"productId": productId,
		})
	})
	// se fossemos chamar essa rota, seria /produtos/:productId
	productsGroup.PUT("/:productId", Update)
	// se fossemos chamar essa rota, seria /produtos/:productId
	productsGroup.DELETE("/:productId", func(c *gin.Context) {})

	server.Run(":8080")
}

// /*
// {
// 	name: "Produto 1"
// }
// */
