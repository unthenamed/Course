package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// Inisialisasi dan  koneksi ke database
const (
	host     = "unthenamed-unthenamed.h.aivencloud.com"
	port     = 13689
	user     = "avnadmin"
	password = "AVNS_DilIFOqbMOlCyx0DADG"
	dbname   = "defaultdb"
	sslmode  = "require"
)

type Student struct {
	id        int
	name      string
	email     string
	address   string
	birthdate time.Time
	gender    string
}

func isNil(err error) (state bool) {
	if err != nil {
		panic(err)
	}
	return
}

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

// fungsi untuk membuka koneksi ke database
func psqlConnect(sqlPath string) (db *sql.DB, err error) {
	if db, err = sql.Open("postgres", sqlPath); !isNil(err) {
		if err = db.Ping(); !isNil(err) {
			fmt.Println("Connected to database!")
		}
	}
	return
}

func main() {
	// koneksi ke ddatabase
	db, _ := psqlConnect(psqlInfo)
	defer db.Close()

	student1 := Student{id: 1, name: "Abdul", email: "abdul@gmail.com", address: "Kp. Nusa", birthdate: time.Date(2002, 8, 20, 0, 0, 0, 0, time.Local), gender: "M"}

	deleteStudent(db, strconv.Itoa(student1.id))
	addStudent(db, student1)
	updateStudent(db, student1)

}

func addStudent(db *sql.DB, st Student) {
	sqlStatment := `INSERT INTO mst_student
					(id, name, email, address, birth_date, gender) 
					VALUES ($1,$2,$3,$4,$5,$6)`

	result, err := db.Exec(sqlStatment, st.id, st.name, st.email, st.address, st.birthdate, st.gender)
	if !isNil(err) {
		fmt.Println("Sukses INSERT ke database", result)
	}
}

func updateStudent(db *sql.DB, st Student) {
	sqlStatment := `UPDATE mst_student
					SET name=$2, email=$3, address=$4, birth_date=$5, gender=$6
					WHERE id=$1`

	result, err := db.Exec(sqlStatment, st.id, st.name, st.email, st.address, st.birthdate, st.gender)
	if !isNil(err) {
		fmt.Println("Sukses UPDATE ke database", result)
	}
}

func deleteStudent(db *sql.DB, id string) {
	sqlStatment := `DELETE FROM mst_student WHERE id=$1`

	result, err := db.Exec(sqlStatment, id)
	row, err := result.RowsAffected()
	if !isNil(err) {
		if row != 0 {
			fmt.Println("Sukses DELETE ke database!")
		} else {
			fmt.Println("Student dengan id", id, "tidak ditemukan!")
		}
	}
}
