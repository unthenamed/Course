package app

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type tCustomers struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

func (t *tCustomers) getById() error {
	statment := `SELECT * FROM customers WHERE id=$1`
	err := DB.QueryRow(statment, t.Id).Scan(&t.Id, &t.Name, &t.PhoneNumber, &t.Address)
	return err
}

/*
    Create Customer
*/

// POST METHOD
func CreateCustomers(c *gin.Context) {

	// Mengambil nilai body
	var nCustomers tCustomers
	err := c.ShouldBind(&nCustomers)
	if err != nil {
		badGatewayRespons(err.Error(), c)
		return
	}

	// Mencari id yang tidak di gunakan
	if nCustomers.Id == 0 {
		search := `SELECT COUNT(id) FROM customers`
		err = DB.QueryRow(search).Scan(&nCustomers.Id)
		if err != nil {
			badGatewayRespons(err.Error(), c)
			return
		}
	}

	// Membuat data baru ke database
	statment := `INSERT INTO customers (id , name, phone, address)
				VALUES ( $1, $2, $3, $4 )`
	_, err = DB.Exec(statment, nCustomers.Id, nCustomers.Name, nCustomers.PhoneNumber, nCustomers.Address)
	if err != nil {
		badGatewayRespons(err.Error(), c)
		return
	}

	// Memberikan respon  created
	createdRespons(nCustomers, c)
}

/*
    Get Customer
*/

// GET METHOD | :id
func GetCustomersById(c *gin.Context) {
	// Mengambil path param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		badRequestRespons("id", c)
		return
	}

	// mengambil data dari database
	matchedCustomers := tCustomers{}
	statment := `SELECT * FROM customers WHERE id = $1`
	err = DB.QueryRow(statment, id).Scan(&matchedCustomers.Id, &matchedCustomers.Name, &matchedCustomers.PhoneNumber, &matchedCustomers.Address)

	// Bedakan antara id tidak ada dan kesalahan database
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower("no rows in result set")) {
			notFoundRespons("id", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	// Memberikan respon OK
	okRespons(matchedCustomers, c)
}

/*
    Update Customer
*/

// PUT METHOD / :id
func UpdateCustomersById(c *gin.Context) {

	// Deklarasi  struct kosong dan error
	var updatedCustomers tCustomers
	var err error

	// Mengambil path param
	updatedCustomers.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		badRequestRespons("id", c)
		return
	}

	// mengambil isi body request
	if err = c.ShouldBind(&updatedCustomers); err != nil {
		badRequestRespons("body", c)
		return
	}

	// Mengambil old data
	var matchedCustomers tCustomers
	statment := `SELECT * FROM customers WHERE id = $1`
	err = DB.QueryRow(statment, updatedCustomers.Id).Scan(&matchedCustomers.Id, &matchedCustomers.Name, &matchedCustomers.PhoneNumber, &matchedCustomers.Address)

	// Bedakan antara id tidak ada dan kesalahan database
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower("no rows in result set")) {
			notFoundRespons("id", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	// parsial update
	if strings.TrimSpace(updatedCustomers.Name) == "" {
		updatedCustomers.Name = matchedCustomers.Name
	}
	if strings.TrimSpace(updatedCustomers.PhoneNumber) == "" {
		updatedCustomers.PhoneNumber = matchedCustomers.PhoneNumber
	}
	if strings.TrimSpace(updatedCustomers.Address) == "" {
		updatedCustomers.Address = matchedCustomers.Address
	}

	// mengupdate database
	statment = `UPDATE customers SET name=$2, phone=$3, address=$4 WHERE id =$1`
	_, err = DB.Exec(statment, updatedCustomers.Id, updatedCustomers.Name, updatedCustomers.PhoneNumber, updatedCustomers.Address)
	if err != nil {
		badGatewayRespons(err.Error(), c)
		return
	}

	// Memberikan respon oke
	okRespons(updatedCustomers, c)
}

/*
    Delete Customer
*/

// DELETE METHOD / :id
func DeleteCustomersById(c *gin.Context) {

	// Mengambil path param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		badRequestRespons("id", c)
		return
	}

	// Ajukan hapus ke database
	statment := `DELETE FROM customers WHERE id = $1`
	result, err := DB.Exec(statment, id)
	if err != nil {
		badGatewayRespons(err.Error(), c)
	}

	// Pengecekan jika id tidak ada
	r, _ := result.RowsAffected()
	if r == 0 {
		notFoundRespons("id", c)
		return
	}

	// Memberikan respon OK
	okRespons("deleted", c)

}
