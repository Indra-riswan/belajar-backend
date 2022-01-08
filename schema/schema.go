package schema

type BooksInput struct {
	Tittle string `json:"tittle" binding:"required"`
	Price  int    `json:"price" binding:"required,number"`
}
