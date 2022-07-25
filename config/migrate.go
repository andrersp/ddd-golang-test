package config

import (
	"sistema/entity"

	"gorm.io/gorm"
)

func DBMigrate() (*gorm.DB, error) {

	conn, err := ConnectDB()

	if err != nil {
		return nil, err
	}
	conn.AutoMigrate(entity.Book{}, entity.User{})

	return conn, nil

}
