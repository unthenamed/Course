package app

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type tEmployees struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

// method membuat data baru
func (t *tEmployees) create() (err error) {
	statment := `INSERT INTO
				employees (name, phone, address) 
				VALUES ($1,$2,$3) RETURNING id`
	err = DB.QueryRow(statment, t.Name, t.PhoneNumber, t.Address).Scan(&t.Id)
	return
}

// method  ngupdate database berdasarkan isi struct
func (t *tEmployees) update() (err error) {
	statment := `UPDATE employees SET name=$2, phone=$3, address=$4 WHERE id=$1`
	_, err = DB.Exec(statment, t.Id, t.Name, t.PhoneNumber, t.Address)
	return
}

// Method menghapus data berdasarkan  id
func (t *tEmployees) delete() error {
	statment := `DELETE FROM employees WHERE id = $1`
	result, err := DB.Exec(statment, t.Id)
	r, _ := result.RowsAffected()
	if r == 0 && err == nil {
		return errors.New("Not Found!")
	}
	return err
}

// Methid mengambil data row berdasarkan id
func (t *tEmployees) read() (err error) {
	statment := `SELECT name, phone, address FROM employees WHERE id = $1`
	err = DB.QueryRow(statment, t.Id).Scan(&t.Name, &t.PhoneNumber, &t.Address)
	return
}

/*
    Create Employee
*/

func CreateEmployees(c *gin.Context) {
	var nEmployees tEmployees

	// Mengambil isi body
	if err := c.ShouldBind(&nEmployees); err != nil {
		badRequestRespons("body", c)
		return
	}

	// Mengirimkan untuk membuat data baru di database
	if err := nEmployees.create(); err != nil {
		badGatewayRespons(err.Error(), c)
		return
	}

	// Memberikam respon oke
	createdRespons(nEmployees, c)

}

/*
    Get Employee
*/

func GetEmployeesById(c *gin.Context) {
	var nEmployees tEmployees
	var err error

	// Mengambil path param
	nEmployees.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		badRequestRespons("id", c)
		return
	}

	// mengisi instance berdasarkann id instance yang ada di database
	if err = nEmployees.read(); err != nil {
		if strContains(err.Error(), "no rows") {
			notFoundRespons("id", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return

	}

	// Memberikan respon hasil data
	okRespons(nEmployees, c)
}

/*
    Update Employee
*/

func UpdateEmployees(c *gin.Context) {
	var nEmployees tEmployees
	var mEmployees tEmployees
	var err error

	// Mengambil path param
	if nEmployees.Id, err = strconv.Atoi(c.Param("id")); err != nil {
		badRequestRespons("id", c)
		return
	}

	// Mengambil isi body
	if err = c.ShouldBind(&nEmployees); err != nil {
		badRequestRespons("body", c)
		return
	}

	// Mengambil data yang belium di ubah
	mEmployees.Id = nEmployees.Id
	if err = mEmployees.read(); err != nil {
		if strContains(err.Error(), "no rows") {
			notFoundRespons("id", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	// jika filed di data baru kosong tgunakan data lama
	if strings.TrimSpace(nEmployees.Name) == "" {
		nEmployees.Name = mEmployees.Name
	}
	if strings.TrimSpace(nEmployees.PhoneNumber) == "" {
		nEmployees.PhoneNumber = mEmployees.PhoneNumber
	}
	if strings.TrimSpace(nEmployees.Address) == "" {
		nEmployees.Address = mEmployees.Address
	}

	// update hasil data ke database
	if err = nEmployees.update(); err != nil {
		badGatewayRespons(err.Error(), c)
		return
	}

	// berikan respom okee
	okRespons(nEmployees, c)
}

/*
    Delete Employee
*/

func DeleteEmployees(c *gin.Context) {
	var nEmployees tEmployees
	var err error

	// Mengambil path param
	if nEmployees.Id, err = strconv.Atoi(c.Param("id")); err != nil {
		badRequestRespons("id", c)
		return
	}

	// Mengirim instance untuk di hapus
	if err = nEmployees.delete(); err != nil {
		if strContains(err.Error(), "Not Found!") {
			notFoundRespons("id", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	// Memberikan respon terhapus
	okRespons("deleted", c)
}
