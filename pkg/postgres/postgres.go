package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func New() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=root dbname=employee sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
