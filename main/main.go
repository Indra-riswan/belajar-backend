package main

import (
	"log"

	"github.com/Indra-riswan/belajar-backend/books"
	"github.com/Indra-riswan/belajar-backend/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/db-belajar-backend?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal Connect Database")
	}
	db.AutoMigrate(&books.Book{})

	booksRepository := books.NewRepository(db)

	// book, err := booksRepository.FindById(1)
	// fmt.Println("Tittle:", book.Tittle)

	// for _, book := range books {
	// 	fmt.Println("Tittle:", book.Tittle)
	// 	fmt.Println("Description:", book.Description)
	// 	fmt.Println("Harga:", book.Ajinne)
	// }

	book := books.Book{
		Tittle:      "Anak Duda",
		Description: "Cerita Anak Duda",
		Penulis:     "Venger",
		Ajinne:      "20000",
	}
	booksRepository.Create(book)

	r := gin.Default()

	v1 := r.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hallo", handler.HalloHandler)
	v1.GET("/books/:id/:tittle", handler.BooksByIdHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.BooksPostHandler)

	r.Run()

}
