package config

import (
	"sistema/entity"

	"gorm.io/gorm"
)

func DBMigrate(db gorm.DB) (err error) {

	err = db.AutoMigrate(entity.Book{}, entity.User{})

	return

}
