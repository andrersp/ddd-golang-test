package service

import (
	"sistema/domain/books/model"
	"sistema/domain/books/repository"
)

type BookService interface {
	GetById(bookID uint) (*model.Book, error)
	Create(book model.Book) (*model.Book, error)
}

type bookService struct {
	repository repository.BookRepository
}

func (b *bookService) GetById(bookID uint) (*model.Book, error) {

	return nil, nil
}

func (b *bookService) Create(book model.Book) (*model.Book, error) {

	b.repository.Create(book)
	return nil, nil
}

func NewBookService(repository repository.BookRepository) BookService {
	return &bookService{
		repository: repository,
	}
}
