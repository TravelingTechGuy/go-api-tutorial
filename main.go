package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//load list of books to memory
	if result, err := loadBooks(); err == nil {
		books = append(books, result...)
	} else {
		log.Fatalf("Exiting: error occured: %s", err)
		os.Exit(-1)
	}
	router := gin.Default()
	setRoutes(router)
	router.Run("localhost:8080")
}
