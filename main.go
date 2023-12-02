package main

import (
	"database/sql"
	// _ "github.com" EROR
	"log"
)

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "rizalfachrezi:password@tcp(localhost:3306)/mkp")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
