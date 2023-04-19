package main

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func openDB() (*sql.DB, error) {
	path := filepath.Join(".", "data", "books.db")
	//delete existing DB - ignore if not there
	_ = os.Remove(path)
	//create new DB file
	db, err := sql.Open("sqlite", path)
	return db, err
}

func createTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS books(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL UNIQUE,
		author TEXT NOT NULL,
		quantity INTEGER NOT NULL
	);
	`
	_, err := db.Exec(query)
	return err
}

func insertData(db *sql.DB) error {
	bs := loadBooksStatic()
	for _, b := range bs {
		_, err := db.Exec("INSERT INTO books(title, author, quantity) values(?,?,?)", b.Title, b.Author, b.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func getAllBooks(db *sql.DB) ([]book, error) {
	rows, err := db.Query("SELECT * FROM books ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var bs []book
	for rows.Next() {
		var b book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity); err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}
	return bs, nil
}

func loadBooksDB() ([]book, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	err = createTable(db)
	if err != nil {
		return nil, err
	}
	err = insertData(db)
	if err != nil {
		return nil, err
	}
	b, err := getAllBooks(db)
	if err != nil {
		return nil, err
	}
	err = db.Close()
	if err != nil {
		return nil, err
	}
	return b, nil
}
