package library

import (
	"database/sql"
	"enigma_laundy/config"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.SqlInfo)
	if !IsNil(err) {
		if err = db.Ping(); !IsNil(err) {
			fmt.Println("Successfully Connected!")
		}
	}
	return db, err
}
