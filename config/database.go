package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config = ConfigDB{}

type ConfigDB struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func ConnectDB() (*gorm.DB, error) {

	config.Read()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		config.Host, config.User, config.Password, config.Dbname, config.Port)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *ConfigDB) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
