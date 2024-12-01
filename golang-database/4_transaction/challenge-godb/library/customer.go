package library

import (
	"database/sql"
	"fmt"
	"time"
)

func CreateCustomer(instance Customer, db *sql.DB) {
	sqlStatment := `
					INSERT INTO customer (customer_id, name,
					phone, address, created_at, update_at)
					VALUES ($1, $2, $3, $4, $5, $6);
					`
	result, err := db.Exec(sqlStatment, instance.Customer_id, instance.Name, instance.Phone, instance.Address, time.Now().In(Location).Format("2006-01-02"), instance.Update_at)

	if !IsNil(err) {
		fmt.Println("New Customer added!", result)
	}

}

func ViewOfListCustomer(db *sql.DB) (instance []Customer) {
	sqlStatment := "SELECT * FROM customer"

	rows, err := db.Query(sqlStatment)

	if !IsNil(err) {
		for rows.Next() {
			stamp := Customer{}
			err = rows.Scan(&stamp.Customer_id, &stamp.Name, &stamp.Phone, &stamp.Address, &stamp.Created_at, &stamp.Update_at)
			instance = append(instance, stamp)
		}

		if err = rows.Err(); !IsNil(err) {

		}
	}

	return
}
func ViewDetailCustomerById(id int, db *sql.DB) (stamp Customer) {
	sqlStatment := "SELECT * FROM customer WHERE customer_id = $1"

	err := db.QueryRow(sqlStatment, id).Scan(&stamp.Customer_id, &stamp.Name, &stamp.Phone, &stamp.Address, &stamp.Created_at, &stamp.Update_at)

	if !IsNilCustom(err, "Customer not found.") {
	}

	return
}

func UpdateCustomer(instance Customer, db *sql.DB) {
	sqlStatment := `
						UPDATE customer
						SET name = $2, phone = $3, 
						address = $4, update_at = $5 
						WHERE customer_id = $1;
	`
	result, err := db.Exec(sqlStatment, instance.Customer_id, instance.Name, instance.Phone, instance.Address, time.Now().In(Location))

	if !IsNil(err) {
		fmt.Println("Customer updated!", result)
	}
}

func DeleteCustomer(id string, db *sql.DB) {
	sqlStatment := "DELETE FROM customer WHERE customer_id = $1"
	result, err := db.Exec(sqlStatment, id)
	if !IsNil(err) {
		fmt.Println("Customer with id", id, "deleted!", result)
	}
}
