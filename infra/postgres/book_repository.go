package postgres

import (
	"errors"
	"sistema/entity"

	"gorm.io/gorm"
)

type dbBookRepository struct {
	db *gorm.DB
}

func (r *dbBookRepository) GetById(bookID uint) (*entity.Book, error) {
	return nil, nil
}

func (r *dbBookRepository) Create(book entity.Book) (*entity.Book, error) {

	r.db.Model(&entity.Book{}).Create(&book)

	return &book, nil
}

func (r *dbBookRepository) GetByTitle(title string) (book *entity.Book, err error) {

	err = r.db.Where("title = ?", title).First(&book).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			book = nil
		}

	}
	return
}

func NewDbBookRepository(db *gorm.DB) *dbBookRepository {
	return &dbBookRepository{
		db: db,
	}
}
