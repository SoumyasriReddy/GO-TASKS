package prod

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/product", getProducts).Methods("GET")
	r.HandleFunc("/product/{ID}", getProduct).Methods("GET")
	r.HandleFunc("/product", createProduct).Methods("POST")
	r.HandleFunc("/product/{ID}", updateProduct).Methods("PUT")
	r.HandleFunc("/product/{ID}", deleteProduct).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
