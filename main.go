package main

import (
	"fmt"
	"log"
	"sistema/config"
	"sistema/domain/books/model"
	book "sistema/domain/books/service"
	user_model "sistema/domain/users/model"
	user "sistema/domain/users/service"
	"sistema/infra/postgres"
)

func main() {
	fmt.Println("start")

	conn, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	config.DBMigrate()

	userHandler := user.NewUserService(postgres.NewDbUserRepository())

	user := user_model.User{
		Name: "Nome Usuario",
		Age:  11,
	}

	userHandler.Create(user)

	bookHandler := book.NewBookService(postgres.NewDbBookRepository(conn))

	books := []model.Book{
		{
			Title: "Titulo do Livro 1",
		},
		{
			Title: "Titulo do Livro 2",
		},
		{
			Title: "Titulo do Livro 4",
		},
		{
			Title: "Titulo do Livro 4",
		},
	}

	for _, book := range books {
		bookHandler.Create(book)
	}

}
