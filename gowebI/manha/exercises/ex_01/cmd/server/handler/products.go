package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/batatinha123/bootcamp-meli-web/internal/products"
	"github.com/gin-gonic/gin"
)

type CreateRequestDto struct {
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Count     int     `json:"count"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
}

type UpdateRequestDto struct {
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Count     int     `json:"count"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
}

type UpdateNameRequestDto struct {
	Name string `json:"name"`
}

type ProductHandler struct {
	service products.Service
}

func NewProduct(p products.Service) *ProductHandler {
	return &ProductHandler{
		service: p,
	}
}

func (c *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		queryName, queryParamNameExists := ctx.GetQuery("name")
		queryPublished, queryParamPublishedExists := ctx.GetQuery("published")

		var filter products.Filter

		if queryParamNameExists {
			filter.Name = queryName
		}

		if queryParamPublishedExists {
			if queryPublished == "true" {
				var published bool = true
				filter.Published = &published
			} else {
				var published bool = false
				filter.Published = &published
			}
		}

		p, err := c.service.GetAll(filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(p) == 0 {
			ctx.Status(http.StatusNoContent)
			return
		}

		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateRequestDto
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		// quando chamamos a service, os dados já estarão tratados
		p, err := c.service.Store(req.Name, req.Code, req.Color, req.Count, req.Price, req.Published)
		if err != nil {
			switch {
			case errors.Is(err, products.ErrProductAlreadyExists):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		ctx.JSON(http.StatusCreated, p)
	}
}

func (c *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// forma de se fazermos uma conversão de alfa númerico para inteiro
		// strconv.Atoi(ctx.Param("id"))
		id, err := strconv.ParseUint(ctx.Param("productId"), 10, 0)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		var req UpdateRequestDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "o nome do produto é obrigatório"})
			return
		}

		if req.Code == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "o código do produto é obrigatório"})
			return
		}

		if req.Count == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "a quantidade é obrigatória"})
			return
		}

		if req.Price == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O preço é obrigatório"})
			return
		}

		p, err := c.service.Update(id, req.Name, req.Code, req.Color, req.Count, req.Price, req.Published)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductHandler) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// product, err = s.repo.FindByID(id)
		// if err != nil {
		//
		// }
		// if request.Name != "" {
		//	product.Name = request.Name
		// }
		//
		// s.repo.Update(id, product)

		// o paramâmetro 0 define para que o GO pegue os bits do processador em que está rodando a aplicação
		id, err := strconv.ParseUint(ctx.Param("productId"), 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		var req UpdateNameRequestDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O nome do produto é obrigatório"})
			return
		}
		p, err := c.service.UpdateName(id, req.Name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("productId"), 10, 0)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		// poderiamos usar o http.StatusNoContent -> 204 tbm!
		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})
	}
}
