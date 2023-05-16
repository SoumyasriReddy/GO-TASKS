package library

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/book", getBooks).Methods("GET")
	r.HandleFunc("/book/{book_id}", getBook).Methods("GET")
	r.HandleFunc("/book", createBooks).Methods("POST")
	r.HandleFunc("/book/{book_id}", updateBooks).Methods("PUT")
	r.HandleFunc("/book/{book_id}", deleteBooks).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
