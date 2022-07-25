package config

import (
	"sistema/domain"

	"gorm.io/gorm"
)

func DBMigrate(db gorm.DB) (err error) {

	err = db.AutoMigrate(domain.Book{}, domain.User{})

	return

}
