package db

import (
	"database/sql"
	//"os"
	//"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:P%40ssw%40rd%2Fkholeur9%2Fpgsql@localhost:5432/dhaclub?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}