package repository

import "sistema/domain/books/model"

type BookRepository interface {
	GetById(bookID uint) (*model.Book, error)
	Create(book model.Book) (*model.Book, error)
}
