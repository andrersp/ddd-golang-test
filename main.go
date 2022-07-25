package main

import (
	"fmt"
	"log"
	"sistema/config"
	"sistema/domain"
	"sistema/services"

	"sistema/infra/postgres"
)

func main() {
	fmt.Println("start")

	conn, err := config.ConnectDBSQLITE()
	if err != nil {
		log.Fatal(err)
	}

	if err := config.DBMigrate(*conn); err != nil {
		log.Fatal(err)
	}

	userHandler := services.NewUserService(postgres.NewDbUserRepository(conn))

	user := domain.User{
		Name: "Nome Usuario",
		Age:  11,
	}

	userHandler.Create(user)

	bookHandler := services.NewBookService(postgres.NewDbBookRepository(conn))

	books := []domain.Book{}

	for i := 1; i < 500; i++ {
		books = append(books, domain.Book{Title: fmt.Sprintf("Titulo livro %d", i)})

	}

	for _, book := range books {
		result, err := bookHandler.Create(book)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Livro Cadastrado: %s\n", result.Title)
		}

	}

	for i := 1; i < 500; i++ {
		book, err := bookHandler.GetById(uint(i))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(*&book.Title)
		}

	}

}
