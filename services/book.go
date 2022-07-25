package services

import (
	"errors"

	"sistema/domain"
	"sistema/domain/repository"
)

type bookService struct {
	repository repository.BookRepository
}

func (b *bookService) GetById(bookID uint) (*domain.Book, error) {

	book, err := b.repository.GetById(bookID)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b *bookService) Create(book domain.Book) (*domain.Book, error) {

	if book, _ := b.repository.GetByTitle(book.Title); book != nil {
		return nil, errors.New("Book Exist")

	}

	_, err := b.repository.Create(book)

	if err != nil {
		return nil, err

	}
	return &book, nil
}

func (b *bookService) GetByTitle(title string) (*domain.Book, error) {
	return nil, nil
}

func NewBookService(repository repository.BookRepository) *bookService {
	return &bookService{
		repository: repository,
	}
}
