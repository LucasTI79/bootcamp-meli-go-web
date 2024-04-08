package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// inicializando uma instancia do gin com os middlewares default
	router := gin.Default()

	// criando uma rota para o método GET
	router.GET("/hello-world", func(c *gin.Context) {
		//O corpo, cabeçalho e método estão contidos no contexto gin.
		body := c.Request.Body
		header := c.Request.Header
		method := c.Request.Method // GET, POST, PUT, DELETE, etc.

		fmt.Println("Eu recebi algo!")
		fmt.Printf("\tMétodo: %s\n", method)
		fmt.Printf("\tConteúdo do cabeçalho:\n")

		for key, value := range header {
			fmt.Printf("\t\t%s -> %s\n", key, value)
		}

		fmt.Printf("\tO body é um io.ReadCloser:(%s), e para trabalhar com ele teremos que leia depois\n", body)
		fmt.Println("¡Yay!")

		// text/plain
		c.String(200, "Eu recebi!") //Res

		// // application/json;
		// c.JSON(200, gin.H{
		// 	"message": "Hello World!",
		// })
	})

	// criando uma rota com grupos

	//Agora podemos atender solicitações para /gophers/, /gophers/get ou /gophers/info de   uma forma mais elegante.
	gopher := router.Group("/gophers")
	{
		gopher.GET("/", func(c *gin.Context) {
			c.String(200, "Gophers!")
		})
		gopher.GET("/get", func(c *gin.Context) {
			c.String(200, " Get gophers!")
		})
		gopher.GET("/info", func(c *gin.Context) {
			c.String(200, "Info gophers!")
		})
	}

	// rodando o servidor na porta 8080
	router.Run(":8080")
}
