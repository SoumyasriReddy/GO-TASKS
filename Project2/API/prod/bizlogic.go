package prod

import (
	"Project2/model"
	"database/sql"
	"fmt"
	"log"
)

func GetAllProducts(db *sql.DB, product []model.Product) []model.Product {
	rows, err := db.Query("SELECT * FROM Product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var products model.Product
		err := rows.Scan(&products.ID, &products.Name, &products.Price)
		if err != nil {
			log.Fatal(err)
		}
		product = append(product, products)
	}
	return product
}

func DeleteProduct(db *sql.DB, ID string) {
	_, err := db.Exec("DELETE FROM Product WHERE ID = ?", ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete Products from db")
}

func GetProduct(db *sql.DB, products []model.Product, ID string) {
	var product model.Product
	err := db.QueryRow("SELECT * FROM Products WHERE ID = ?", product.ID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateProduct(db *sql.DB, products model.Product) {
	_, err := db.Exec("INSERT INTO Product (ID,Name,Price) VALUES (?, ?, ?)",
		products.ID, products.Name, products.Price)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateProduct(db *sql.DB, products model.Product) {
	_, err := db.Exec("UPDATE Product SET ID=? , Name=?, Price=? WHERE ID=?",
		products.ID, products.Name, products.Price, products.ID)
	if err != nil {
		log.Fatal(err)
		fmt.Println("----In db Err--------")
	}
	fmt.Println("------ HIIIIII Working---------")
}
