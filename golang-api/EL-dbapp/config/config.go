package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     string = "unthenamed-unthenamed.h.aivencloud.com"
	port     int    = 13689
	user     string = "avnadmin"
	password string = "AVNS_DilIFOqbMOlCyx0DADG"
	database string = "el_db"
	cert     string = "ca.pm"
)

var sqlInfo string = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require", user, password, host, port, database)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}
