package main

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func GetDatabase() {
	connStr := "user=root dbname=root sslmode=verify-full"
	db, err := sql.Open("key_share", connStr)

	if err != nil {
		log.Println("error while connecting to database")
	}

	DB = db
}
