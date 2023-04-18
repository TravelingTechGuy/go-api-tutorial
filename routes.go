package main

import (
	"errors"
	"net/http"

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

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id parameter missing"})
		return
	}
	book, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if book.Quantity == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no more copies left of book " + id})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id parameter missing"})
		return
	}
	book, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	if b, _ := bookById(newBook.ID); b != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "book with id " + newBook.ID + " already exists"})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func setRoutes(router *gin.Engine) {
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PATCH("/books/checkout", checkoutBook)
	router.PATCH("/books/return", returnBook)
}
