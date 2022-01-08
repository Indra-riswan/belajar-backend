package books

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(books Book) (Book, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (n *repository) FindAll() ([]Book, error) {
	var books []Book

	err := n.db.Find(&books).Error

	return books, err

}
func (n *repository) FindById(ID int) (Book, error) {
	var book Book

	err := n.db.Take(&book, ID).Error

	return book, err
}

func (n *repository) Create(books Book) (Book, error) {
	err := n.db.Create(&books).Error

	return books, err
}
