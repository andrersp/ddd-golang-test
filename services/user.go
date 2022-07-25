package services

import (
	"errors"

	"sistema/domain/user"

	"sistema/entity"
)

type UserService struct {
	repository user.UserRespository
}

func (u *UserService) Create(user entity.User) (*entity.User, error) {

	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	return u.repository.Create(user)
}

func (u *UserService) GetById(userID uint) (user *entity.User, err error) {

	user, err = u.repository.GetById(userID)

	if err != nil {
		err = errors.New("user not found")
	}

	return
}

func NewUserService(repository user.UserRespository) *UserService {
	return &UserService{
		repository: repository,
	}
}
