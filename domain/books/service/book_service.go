package service

import (
	"errors"
	"sistema/domain/books/repository"
	"sistema/entity"
)

type BookService interface {
	GetById(bookID uint) (*entity.Book, error)
	Create(book entity.Book) (*entity.Book, error)
	GetByTitle(title string) (*entity.Book, error)
}

type bookService struct {
	repository repository.BookRepository
}

func (b *bookService) GetById(bookID uint) (*entity.Book, error) {

	return nil, nil
}

func (b *bookService) Create(book entity.Book) (*entity.Book, error) {

	if book, _ := b.repository.GetByTitle(book.Title); book != nil {
		return nil, errors.New("Book Exist")

	}

	_, err := b.repository.Create(book)

	if err != nil {
		return nil, err

	}
	return &book, nil
}

func (b *bookService) GetByTitle(title string) (*entity.Book, error) {
	return nil, nil
}

func NewBookService(repository repository.BookRepository) BookService {
	return &bookService{
		repository: repository,
	}
}
