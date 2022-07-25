package service

import (
	"errors"
	"sistema/domain/users/repository"
	"sistema/entity"
)

type UserService interface {
	GetById(userID uint) (*entity.User, error)
	Create(user entity.User) (*entity.User, error)
}

type userService struct {
	repository repository.UserRespository
}

func (u *userService) Create(user entity.User) (*entity.User, error) {

	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	return u.repository.Create(user)
}

func (u *userService) GetById(userID uint) (user *entity.User, err error) {

	user, err = u.repository.GetById(userID)

	if err != nil {
		err = errors.New("user not found")
	}

	return
}

func NewUserService(repository repository.UserRespository) UserService {
	return &userService{
		repository: repository,
	}
}
