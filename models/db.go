package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "dbname=lemonade user=postgres password=secret sslmode=disable")

	log.Print(db)
	log.Print(err)
	return db, err
}
