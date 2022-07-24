package postgres

import (
	"fmt"
	"sistema/domain/users/model"
)

type DBUserRepository interface {
	GetById(userID uint) (*model.User, error)
	Create(user model.User) (*model.User, error)
}

type dbUserRepository struct {
}

func NewDbUserRepository() DBUserRepository {
	return &dbUserRepository{}
}

func (r *dbUserRepository) Create(user model.User) (*model.User, error) {

	fmt.Println(user)

	return nil, nil
}

func (r *dbUserRepository) GetById(userID uint) (*model.User, error) {
	return nil, nil
}
