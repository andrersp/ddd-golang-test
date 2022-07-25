package services

import (
	"errors"

	"sistema/domain/book"
	"sistema/entity"
)

type bookService struct {
	repository book.BookRepository
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

func NewBookService(repository book.BookRepository) *bookService {
	return &bookService{
		repository: repository,
	}
}
