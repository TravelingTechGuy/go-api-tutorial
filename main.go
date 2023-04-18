package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error reading env file: %s", err)
		os.Exit(1)
	}

	src := os.Getenv("BOOKS_SOURCE")
	port := os.Getenv("PORT")
	log.Printf("Books will be read from %s", src)
	//load list of books to memory
	if result, err := loadBooks(src); err == nil {
		books = append(books, result...)
	} else {
		log.Fatalf("Exiting: error occured: %s", err)
		os.Exit(2)
	}
	router := gin.Default()
	setRoutes(router)
	router.Run("localhost:" + port)
}
