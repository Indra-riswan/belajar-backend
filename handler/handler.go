package handler

import (
	"fmt"
	"net/http"

	"github.com/Indra-riswan/belajar-backend/schema"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"First page said": "selamat datang di halaman utama",
	})
}
func HalloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"hallo pages said": "hallo-Indra",
	})
}

func BooksByIdHandler(c *gin.Context) {
	id := c.Param("id")
	tittle := c.Param("tittle")

	c.JSON(200, gin.H{
		"id":     id,
		"tittle": tittle,
	})
}

func QueryHandler(c *gin.Context) {
	tittle := c.Query("tittle")
	price := c.Query("price")
	c.JSON(200, gin.H{
		"tittle": tittle,
		"price":  price,
	})
}

func BooksPostHandler(c *gin.Context) {
	var booksinput schema.BooksInput

	err := c.ShouldBindJSON(&booksinput)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Erorr on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMassage)
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"tittle": booksinput.Tittle,
		"price":  booksinput.Price,
	})
}
