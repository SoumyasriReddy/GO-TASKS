package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	// Connect to the database
	db, err := sql.Open("mysql", "root:Nanna@143@tcp(127.0.0.1:3306)/library_DB")
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to database!")
	return db
}
