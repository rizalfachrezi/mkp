package main

import (
	"database/sql"
	// _ "github.com/lib/pq" EROR
	"log"
)

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://rizalfachrezi:password@localhost/mkp?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
