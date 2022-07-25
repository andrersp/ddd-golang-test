package repository

import "sistema/domain"

type BookRepository interface {
	GetById(bookID uint) (*domain.Book, error)
	Create(book domain.Book) (*domain.Book, error)
	GetByTitle(title string) (*domain.Book, error)
}
