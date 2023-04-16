package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadBooksStatic() {
	var staticBooks = []book {
		{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
		{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
		{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
	}

	books = append(books, staticBooks...)
}

func loadBooksJSON() {
	//Read the file
	content, err := os.ReadFile("./data/books.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	// Now let's unmarshall the data into `payload`
	var payload []book
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	
	//Load the books to our memor structure
	books = append(books, payload...)
}


func loadBooks() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	src := os.Getenv("BOOKS_SOURCE")
	switch src {
		case "static":
			loadBooksStatic()
		case "json", "db":
			loadBooksJSON()
		default:
			log.Fatal("wrong books source specified in env file")
	}
	log.Printf("Books read from %s", src)
}
