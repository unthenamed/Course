package library

import (
	"database/sql"
	"fmt"
	"time"
)

func CreateService(instance Service, db *sql.DB) {
	sqlStatment := `
					INSERT INTO service (service_id, service_name,
					unit, price, created_at, update_at)
					VALUES ($1, $2, $3, $4, $5, $6);
					`
	result, err := db.Exec(sqlStatment, instance.Id, instance.Name, instance.Unit, instance.Price, time.Now().In(Location).Format("2006-01-02"), instance.Update_at)

	if !IsNil(err) {
		fmt.Println("New Service added!", result)
	}

}

func ViewOfListService(db *sql.DB) (instance []Service) {
	sqlStatment := "SELECT * FROM service"

	rows, err := db.Query(sqlStatment)

	if !IsNil(err) {
		for rows.Next() {
			stamp := Service{}
			err = rows.Scan(&stamp.Id, &stamp.Name, &stamp.Unit, &stamp.Price, &stamp.Created_at, &stamp.Update_at)
			instance = append(instance, stamp)
		}

		if err = rows.Err(); !IsNil(err) {

		}
	}

	return
}
func ViewDetailServiceById(id int, db *sql.DB) (stamp Service) {
	sqlStatment := "SELECT * FROM service WHERE service_id = $1"

	err := db.QueryRow(sqlStatment, id).Scan(&stamp.Id, &stamp.Name, &stamp.Unit, &stamp.Price, &stamp.Created_at, &stamp.Update_at)

	if !IsNilCustom(err, "Service not found.") {
	}

	return
}

func UpdateService(instance Service, db *sql.DB) {
	sqlStatment := `
						UPDATE service
						SET service_name = $2, unit = $3, 
						price = $4, update_at = $5 
						WHERE service_id = $1;
	`
	result, err := db.Exec(sqlStatment, instance.Id, instance.Name, instance.Unit, instance.Price, time.Now().In(Location).Format("2006-01-02"))

	if !IsNil(err) {
		fmt.Println("Service updated!", result)
	}
}

func DeleteService(id int, db *sql.DB) {
	sqlStatment := "DELETE FROM service WHERE service_id = $1"
	result, err := db.Exec(sqlStatment, id)
	if !IsNil(err) {
		fmt.Println("Service with id", id, "deleted!", result)
	}
}
