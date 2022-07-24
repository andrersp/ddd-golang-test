package repository

import "sistema/domain/users/model"

type UserRespository interface {
	GetById(userID uint) (*model.User, error)
	Create(user model.User) (*model.User, error)
}
