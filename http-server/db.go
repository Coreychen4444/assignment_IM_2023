package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	// Initialize the database connection here
	// Set up the connection to my database
	connectionString := "corey:123456@tcp(localhost:3306)/mydatabase
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Check the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")
}
