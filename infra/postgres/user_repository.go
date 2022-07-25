package postgres

import (
	"sistema/entity"

	"gorm.io/gorm"
)

type DBUserRepository interface {
	GetById(userID uint) (*entity.User, error)
	Create(user entity.User) (*entity.User, error)
}

type dbUserRepository struct {
	db *gorm.DB
}

func NewDbUserRepository(db *gorm.DB) DBUserRepository {
	return &dbUserRepository{
		db: db,
	}
}

func (r *dbUserRepository) Create(user entity.User) (*entity.User, error) {

	r.db.Model(&entity.User{}).Create(&user)

	// fmt.Println(user)

	return nil, nil
}

func (r *dbUserRepository) GetById(userID uint) (*entity.User, error) {
	return nil, nil
}
