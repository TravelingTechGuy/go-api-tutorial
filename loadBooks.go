package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func loadBooksStatic() []book {
	return []book{
		{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
		{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
		{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
	}
}

func loadBooksJSON() ([]book, error) {
	//Read the file
	content, err := os.ReadFile("./data/books.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return nil, err
	}
	// Now let's unmarshall the data into `payload`
	var payload []book
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return nil, err
	}

	return payload, nil
}

func loadBooks(src string) ([]book, error) {
	switch src {
	case "static":
		return loadBooksStatic(), nil
	case "json", "db":
		return loadBooksJSON()
	default:
		return nil, errors.New("wrong books source specified in env file")
	}
}
