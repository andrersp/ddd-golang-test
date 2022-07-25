package repository

import "sistema/domain"

type UserRespository interface {
	GetById(userID uint) (*domain.User, error)
	Create(user domain.User) (*domain.User, error)
}
