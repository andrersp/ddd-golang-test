package postgres

import (
	"sistema/domain/books/model"

	"gorm.io/gorm"
)

type DBBookRepository interface {
	GetById(bookID uint) (*model.Book, error)
	Create(book model.Book) (*model.Book, error)
}

type dbBookRepository struct {
	db *gorm.DB
}

func (r *dbBookRepository) GetById(bookID uint) (*model.Book, error) {
	return nil, nil
}

func (r *dbBookRepository) Create(book model.Book) (*model.Book, error) {

	r.db.Model(&model.Book{}).Create(&book)

	return &book, nil
}

func NewDbBookRepository(db *gorm.DB) DBBookRepository {
	return &dbBookRepository{
		db: db,
	}
}
