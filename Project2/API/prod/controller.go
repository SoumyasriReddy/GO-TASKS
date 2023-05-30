package prod

import (
	"Project2/db/database"
	"Project2/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var product []model.Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.Database()
	product = GetAllProducts(db, product)
	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a product")
	w.Header().Set("Content-Type", "applicattion/json")
	p := mux.Vars(r)
	for index, item := range product {
		if item.ID == p["id"] {
			product = append(product[:index], product[index+1:]...)
			db := database.Database()
			DeleteProduct(db, item.ID)
			break
		}
	}
	json.NewEncoder(w).Encode(product)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for _, item := range product {
		if item.ID == p["id"] {
			db := database.Database()
			GetProduct(db, product, item.ID)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products model.Product
	_ = json.NewDecoder(r.Body).Decode(&products)
	db := database.Database()
	CreateProduct(db, products)
	product = append(product, products)
	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for index, item := range product {
		fmt.Println("-----------Check if it working-------", item.ID)
		if item.ID == p["id"] {
			fmt.Println("-------------Id found-----------")
			product = append(product[:index], product[index+1:]...)
			var updatedProduct model.Product
			_ = json.NewDecoder(r.Body).Decode(&updatedProduct)
			fmt.Println(updatedProduct)
			db := database.Database()
			UpdateProduct(db, updatedProduct)
			fmt.Println("------It is Working-------")
			product = append(product, updatedProduct)
			fmt.Println(product)
			json.NewEncoder(w).Encode(updatedProduct)
		}
	}
}
