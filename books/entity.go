package books

import "time"

type Book struct {
	ID          int
	Tittle      string
	Description string
	Penulis     string
	Ajinne      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
