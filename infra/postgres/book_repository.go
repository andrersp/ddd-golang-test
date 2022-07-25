package postgres

import (
	"errors"
	"sistema/domain"
	"sistema/domain/repository"

	"gorm.io/gorm"
)

type dbBookRepository struct {
	db *gorm.DB
}

func (r *dbBookRepository) GetById(bookID uint) (book *domain.Book, err error) {

	err = r.db.First(&book, bookID).Error
	return
}

func (r *dbBookRepository) Create(book domain.Book) (*domain.Book, error) {

	r.db.Model(&domain.Book{}).Create(&book)

	return &book, nil
}

func (r *dbBookRepository) GetByTitle(title string) (book *domain.Book, err error) {

	err = r.db.Where("title = ?", title).First(&book).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			book = nil
		}

	}
	return
}

func NewDbBookRepository(db *gorm.DB) repository.BookRepository {
	return &dbBookRepository{
		db: db,
	}
}
