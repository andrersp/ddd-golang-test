package main

import (
	"fmt"
	"log"
	"sistema/config"
	book "sistema/domain/books/service"
	user "sistema/domain/users/service"
	"sistema/entity"
	"sistema/infra/postgres"
)

func main() {
	fmt.Println("start")

	conn, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	config.DBMigrate()

	userHandler := user.NewUserService(postgres.NewDbUserRepository(conn))

	user := entity.User{
		Name: "Nome Usuario",
		Age:  11,
	}

	userHandler.Create(user)

	bookHandler := book.NewBookService(postgres.NewDbBookRepository(conn))

	books := []entity.Book{}

	for i := 1; i < 500; i++ {
		books = append(books, entity.Book{Title: fmt.Sprintf("Titulo livro %d", i)})

	}

	for _, book := range books {
		result, err := bookHandler.Create(book)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Livro Cadastrado: %s", result.Title)
		}

	}

}
