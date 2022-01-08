package main

import (
	"github.com/Indra-riswan/belajar-backend/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	v1 := r.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hallo", handler.HalloHandler)
	v1.GET("/books/:id/:tittle", handler.BooksByIdHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.BooksPostHandler)

	r.Run()
}
