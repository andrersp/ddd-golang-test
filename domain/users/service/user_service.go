package service

import (
	"errors"
	"sistema/domain/users/model"
	"sistema/domain/users/repository"
)

type UserService interface {
	GetById(userID uint) (*model.User, error)
	Create(user model.User) (*model.User, error)
}

type userService struct {
	repository repository.UserRespository
}

func (u *userService) Create(user model.User) (*model.User, error) {

	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	return u.repository.Create(user)
}

func (u *userService) GetById(userID uint) (user *model.User, err error) {

	user, err = u.repository.GetById(userID)

	if err != nil {
		err = errors.New("user not found")
	}

	return
}

func NewService(repository repository.UserRespository) UserService {
	return &userService{
		repository: repository,
	}
}
