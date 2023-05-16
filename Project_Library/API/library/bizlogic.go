package library

import (
	"Project_library/model"
	"database/sql"
	"fmt"
	"log"
)

func GetAllBooks(db *sql.DB, book []model.Books) []model.Books {
	rows, err := db.Query("SELECT * FROM Books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var books model.Books
		err := rows.Scan(&books.Book_id, &books.Year_published, &books.Book_title, &books.Author_name, &books.Book_category)
		if err != nil {
			log.Fatal(err)
		}
		book = append(book, books)
	}
	return book
}

func DeleteBook(db *sql.DB, Book_id string) {
	_, err := db.Exec("DELETE FROM Books WHERE Book_id = ?", Book_id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete books db")
}

func GetBook(db *sql.DB, books []model.Books, Book_id string) {
	var book model.Books
	err := db.QueryRow("SELECT * FROM Books WHERE Book_id = ?", book.Book_id).Scan(
		&book.Book_id,
		&book.Year_published,
		&book.Book_title,
		&book.Author_name,
		&book.Book_category,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateBook(db *sql.DB, books model.Books) {
	_, err := db.Exec("INSERT INTO Books (Book_id, Year_published, Book_title, Author_name, Book_category) VALUES (?, ?, ?, ?, ?)",
		books.Book_id, books.Year_published, books.Book_title, books.Author_name, books.Book_category)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateBook(db *sql.DB, books model.Books) {
	_, err := db.Exec("UPDATE Books SET Book_id=? , Year_published=?, Book_title=?, Author_name=?, Book_category=? WHERE Book_id=?",
		books.Book_id, books.Year_published, books.Book_title, books.Author_name, books.Book_category, books.Book_id)
	if err != nil {
		log.Fatal(err)
		fmt.Println("----In db Err--------")
	}
	fmt.Println("------ HIIIIII Working---------")
}
