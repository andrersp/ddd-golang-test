package repository

import (
	"sistema/entity"
)

type BookRepository interface {
	GetById(bookID uint) (*entity.Book, error)
	Create(book entity.Book) (*entity.Book, error)
	GetByTitle(title string) (*entity.Book, error)
}
