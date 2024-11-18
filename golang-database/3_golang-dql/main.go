package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "unthenamed-unthenamed.h.aivencloud.com"
	port     = 13689
	user     = "avnadmin"
	password = "AVNS_DilIFOqbMOlCyx0DADG"
	dbname   = "defaultdb"
	sslmode  = "require"
)

func handlePanic(state *bool) {
	if r := recover(); r != nil {
		fmt.Println(r)
		*state = true
	}
}
func isNil(err error) (result bool) {
	defer handlePanic(&result)
	if err != nil {
		panic(err)
	}
	return
}

func connectDb() *sql.DB {
	sqlURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", sqlURI)
	if !isNil(err) {
		if err = db.Ping(); !isNil(err) {
			fmt.Println("DB Connected!")
		}
	}
	return db
}

type Student struct {
	id        int
	name      string
	email     string
	address   string
	birthdate time.Time
	gender    string
}

func main() {
	db := connectDb()
	defer db.Close()

	allSt := getAllStudent(db)
	for _, st := range allSt {
		fmt.Println(st.id, st.name, st.email, st.address, st.birthdate, st.gender)
	}

	st := getStudentById(db, 3)
	fmt.Println(st.id, st.name, st.email, st.address, st.birthdate, st.gender)

	sch := searchBy(db, "in", "1998-03-14")
	for _, st := range sch {
		fmt.Println(st.id, st.name, st.email, st.address, st.birthdate, st.gender)

	}
}

func scanRows(rows *sql.Rows) (result []Student) {
	for rows.Next() {
		emptyStudent := Student{}
		err := rows.Scan(&emptyStudent.id, &emptyStudent.name, &emptyStudent.email, &emptyStudent.address, &emptyStudent.birthdate, &emptyStudent.gender)
		if !isNil(err) {
			result = append(result, emptyStudent)
		}
	}
	if err := rows.Err(); !isNil(err) {
		fmt.Println(" scan not error")
	}
	return
}

func getAllStudent(db *sql.DB) []Student {
	sqlStatment := "SELECT * FROM mst_student;"

	rows, err := db.Query(sqlStatment)
	if !isNil(err) {
		fmt.Println(" query not error")
	}
	return scanRows(rows)
}

func getStudentById(db *sql.DB, id int) (result Student) {
	sqlStatment := "SELECT * FROM mst_student WHERE id = $1;"

	err := db.QueryRow(sqlStatment, id).Scan(&result.id, &result.name, &result.email, &result.address, &result.birthdate, &result.gender)
	if !isNil(err) {
		fmt.Println("getbyid nor error")
	}
	return
}

func searchBy(db *sql.DB, name string, brithDate string) []Student {
	sqlStatment := "SELECT * FROM mst_student WHERE name LIKE $1 AND birth_date = $2"
	rows, err := db.Query(sqlStatment, "%"+name+"%", brithDate)
	if !isNil(err) {
		fmt.Println("Search OK")
	}
	return scanRows(rows)
}
