package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	// Uma tag struct é fechada com caracteres backtick `
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Id       string `form:"id" json:"id"`
	Active   bool   `form:"active" json:"active" binding:"required"`
}

var employees = map[string]string{
	"644": "Employee A",
	"755": "Employee B",
	"777": "Employee C",
}

func RootPage(ctxt *gin.Context) {
	// text/plain
	ctxt.String(200, "Bem-vindo à empresa Gopher!")
}

// Este manipulador verificará se o id passado pelo cliente existe em nosso banco de dados.
func SearchEmployee(ctxt *gin.Context) {
	// recuperando o valor do path param enviado na request, caso não exista o path param informado, é retornada uma string vazia
	id := ctxt.Param("id")
	fmt.Printf("O id fornecido no path param é: %s\n", id)

	queryParam := ctxt.Query("name")
	fmt.Printf("O valor do query param \"name\" é: %s\n", queryParam)

	employee, ok := employees[id]
	if ok {
		ctxt.String(200, "informação do empregado %s, nome: %s", id, employee)
	} else {
		// notem que estamos enviando o status 404 -> not found para o cliente
		ctxt.String(404, "informação do empregado não existe!")
	}
}

func UpdateEmployee(ctx *gin.Context) {

	var employee Employee
	// o método ShouldBindJSON do nosso contexto irá vincular o conteúdo do corpo
	// para os campos da estrutura que fornecemos
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if employee.Name != "user1" || employee.Password != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "não está autorizado"})
		return
	}

	// aqui teria a lógica para atualizar o empregado na camada de persistência

	ctx.JSON(http.StatusOK, gin.H{"status": "autorizado"})
}

// usamos os path params para identificar recursos de forma única

// usamos os query params para filtrar, ordenar, paginar, etc.
// ?page=1&pageSize=20&sortBy=name&sortDirection=desc&name=monitor

func main() {
	server := gin.Default()
	server.GET("/", RootPage)
	// aqui estamos passando um path param chamado id
	server.GET("/employees/:id", SearchEmployee)
	server.PUT("/employees/:id", UpdateEmployee)
	server.Run(":8080")
}
