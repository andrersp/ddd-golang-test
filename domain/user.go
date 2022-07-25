package domain

import "errors"

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func (u *User) ValidateUser() (err error) {
	if u.Age < 0 {
		err = errors.New("Age menor que 0")
	}
	return
}
