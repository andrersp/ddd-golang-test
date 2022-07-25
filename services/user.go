package services

import (
	"errors"

	"sistema/domain"
	"sistema/domain/repository"
)

type UserService struct {
	repository repository.UserRespository
}

func (u *UserService) Create(user domain.User) (*domain.User, error) {

	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	return u.repository.Create(user)
}

func (u *UserService) GetById(userID uint) (user *domain.User, err error) {

	user, err = u.repository.GetById(userID)

	if err != nil {
		err = errors.New("user not found")
	}

	return
}

func NewUserService(repository repository.UserRespository) *UserService {
	return &UserService{
		repository: repository,
	}
}
