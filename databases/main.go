package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Books struct {
	title  string
	author string
}

func main() {
	db := NewDB()

	http.HandleFunc("/show", SelectBooks(db))
	http.HandleFunc("/insert", InsertBooks(db))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func SelectBooks(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var title, author string
		err := db.QueryRow("select title, author from books").Scan(&title, &author)
		if err != nil {
			panic(err)
		}
		book := &Books{title: title, author: author}
		fmt.Fprintf(rw, "The first book is '%s' by '%s'", book.title, book.author)
	})
}

func InsertBooks(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		_, err := db.Exec(
			"INSERT INTO books (title, author) VALUES (?, ?)",
			"The gopher book",
			"Alvaro Salazar",
		)
		if err != nil {
			panic(err)
		}
	})
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "db_test")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists books(title text, author text)")
	if err != nil {
		panic(err)
	}

	return db
}
