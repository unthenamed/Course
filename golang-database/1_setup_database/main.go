package main

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/lib/pq"
)

const (
	host     string = "unthenamed-unthenamed.h.aivencloud.com"
	port     int    = 13689
	user     string = "avnadmin"
	password string = "AVNS_DilIFOqbMOlCyx0DADG"
	database string = "defaultdb"
	cert     string = "ca.pm"
)

var sqlInfo string = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require", user, password, host, port, database)

func dbConnect() *sql.DB {
	connection, _ := url.Parse(sqlInfo)
	connection.RawQuery = "sslmode=verify-ca;sslrootcert=" + cert

	// Connect to database
	db, err := sql.Open("postgres", connection.String())
	if err != nil {
		panic(err)
	}

	// Ping to database
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected")
	}

	return db
}

func main() {
	db := dbConnect()

	db.Close()

}
