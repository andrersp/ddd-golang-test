package repository

import (
	"sistema/entity"
)

type UserRespository interface {
	GetById(userID uint) (*entity.User, error)
	Create(user entity.User) (*entity.User, error)
}
