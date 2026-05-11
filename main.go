package main

import (
	"database/sql"
	"fmt"
	"golang-basic-api/book/infrastructure"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Connect SQL
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// Ensure DB tables are created
	q := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	) 
	`
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Inject dependencies
	bookStore := infrastructure.NewBookStore(db)
	bookService := infrastructure.NewBookService(bookStore)
	handler := infrastructure.NewEndpointHandler(bookService)

	// Configure routes
	http.HandleFunc("/books", handler.HandleBooks)
	http.HandleFunc("/books/", handler.HandleBook)

	fmt.Println("Server running in http://localhost:8080")

	// Run
	log.Fatal(http.ListenAndServe(":8080", nil))
}
