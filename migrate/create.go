package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres host=localhost sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}
}
