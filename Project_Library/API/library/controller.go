package library

import (
	"Project_Library/db/database"
	"Project_library/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var book []model.Books

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.Database()
	book = GetAllBooks(db, book)
	json.NewEncoder(w).Encode(book)
}

func deleteBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a book")
	w.Header().Set("Content-Type", "applicattion/json")
	p := mux.Vars(r)
	for index, item := range book {
		if item.Book_id == p["book_id"] {
			book = append(book[:index], book[index+1:]...)
			db := database.Database()
			DeleteBook(db, item.Book_id)
			break
		}
	}
	json.NewEncoder(w).Encode(book)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for _, item := range book {
		if item.Book_id == p["book_id"] {
			db := database.Database()
			GetBook(db, book, item.Book_id)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books model.Books
	_ = json.NewDecoder(r.Body).Decode(&books)
	db := database.Database()
	CreateBook(db, books)
	book = append(book, books)
	json.NewEncoder(w).Encode(book)
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for index, item := range book {
		fmt.Println("-----------Check if it working-------", item.Book_id)
		if item.Book_id == p["book_id"] {
			fmt.Println("-------------Id found-----------")
			book = append(book[:index], book[index+1:]...)
			var updatedBook model.Books
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			fmt.Println(updatedBook)
			db := database.Database()
			UpdateBook(db, updatedBook)
			fmt.Println("------It is Working-------")
			book = append(book, updatedBook)
			fmt.Println(book)
			json.NewEncoder(w).Encode(updatedBook)
		}
	}
}
