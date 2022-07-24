package config

import (
	book_model "sistema/domain/books/model"
	user_model "sistema/domain/users/model"

	"gorm.io/gorm"
)

func DBMigrate() (*gorm.DB, error) {

	conn, err := ConnectDB()

	if err != nil {
		return nil, err
	}
	conn.AutoMigrate(book_model.Book{}, user_model.User{})

	return conn, nil

}
