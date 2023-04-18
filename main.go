package main

import (
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var books []book

func bookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book with id " + id + " not found")
}

func main() {
	//load list of books to memory
	if result, err := loadBooks(); err == nil {
		books = append(books, result...)
	} else {
		log.Fatal(err)
		os.Exit(-1)
	}
	router := gin.Default()
	setRoutes(router)
	router.Run("localhost:8080")
}
