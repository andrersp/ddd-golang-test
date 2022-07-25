package postgres

import (
	"sistema/domain"
	"sistema/domain/repository"

	"gorm.io/gorm"
)

type DBUserRepository interface {
	GetById(userID uint) (*domain.User, error)
	Create(user domain.User) (*domain.User, error)
}

type dbUserRepository struct {
	db *gorm.DB
}

func NewDbUserRepository(db *gorm.DB) repository.UserRespository {
	return &dbUserRepository{
		db: db,
	}
}

func (r *dbUserRepository) Create(user domain.User) (*domain.User, error) {

	r.db.Model(&domain.User{}).Create(&user)

	// fmt.Println(user)

	return nil, nil
}

func (r *dbUserRepository) GetById(userID uint) (*domain.User, error) {
	return nil, nil
}
