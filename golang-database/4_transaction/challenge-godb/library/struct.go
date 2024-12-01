package library

import (
	"database/sql"
	"fmt"
	"time"
)

type Customer struct {
	Customer_id int
	Name        string
	Phone       string
	Address     string
	Created_at  time.Time
	Update_at   time.Time
}

type Service struct {
	Id         int
	Name       string
	Unit       string
	Price      int
	Created_at time.Time
	Update_at  time.Time
}

type Order struct {
	order_id        int
	customer_id     int
	order_date      time.Time
	completion_date time.Time
	received_by     string
	created_at      time.Time
	update_at       time.Time
}

type Order_detail struct {
	order_detail_id int
	order_id        int
	service_id      int
	qty             int
}

func (instance *Order) NewOrder(tx *sql.Tx) bool {
	instance.order_id = InputNumber("Input ID Order :")
	if checkOrderId(tx, instance.order_id, "Order ID already exists. Please enter a different ID.") {
		return true
	}
	instance.customer_id = InputNumber("Input ID Customer :")
	if checkCustomerId(tx, instance.customer_id, "Customer not found") {
		return true
	}
	instance.received_by = InputText("Received By :")
	dateStr := InputText("Completion date [YYYY-MM-DD] :")
	instance.order_date = time.Now().In(Location)
	instance.created_at = time.Now().In(Location)
	instance.completion_date, _ = time.Parse("2006-01-02", dateStr)
	return false
}

func (instance *Order_detail) NewDetail(tx *sql.Tx, order Order, autoIncrement int) bool {
	instance.order_detail_id = autoIncrement
	instance.order_id = order.order_id
	instance.service_id = InputNumber("Input ID Service:")
	if checkCustomerId(tx, instance.service_id, "Service not found") {
		return true
	}
	instance.qty = InputNumber("Quantity :")
	return false
}
func checkOrderId(tx *sql.Tx, input int, label string) (status bool) {
	statment := `SELECT order_id FROM "order" WHERE order_id = $1`
	id := 0
	tx.QueryRow(statment, input).Scan(&id)

	if id != 0 {
		status = true
		fmt.Println(label)

	}
	return
}

func checkCustomerId(tx *sql.Tx, input int, label string) (status bool) {
	statment := `SELECT customer_id FROM customer WHERE customer_id = $1`
	id := 0
	tx.QueryRow(statment, input).Scan(&id)

	if id == 0 {
		status = true
		fmt.Println(label)
	}
	return
}

func checkServiceId(tx *sql.Tx, input int, label string) (status bool) {
	statment := `SELECT service_id FROM service WHERE service_id = $1`
	id := 0
	tx.QueryRow(statment, input).Scan(&id)

	if id == 0 {
		status = true
		fmt.Println(label)
	}
	return
}

func getTxId(tx *sql.Tx) (id int) {
	statment := "SELECT COUNT(order_detail_id) FROM order_detail"

	if !validateTransaction(tx.QueryRow(statment).Scan(&id), tx) {
	}
	return
}

func OrderInstance(tx *sql.Tx) (order Order, detail Order_detail, status bool) {
	if order.NewOrder(tx) {
		status = true
		return
	}
	if detail.NewDetail(tx, order, getTxId(tx)+1) {
		status = true
	}
	return
}

func (o *Order) insertOrder(tx *sql.Tx) {
	statment := `INSERT INTO "order" (order_id, customer_id,order_date, completion_date,received_by, created_at, update_at) VALUES ($1,$2,$3,$4,$5,$6,$7);`

	_, err := tx.Exec(statment, o.order_id, o.customer_id, o.order_date, o.completion_date, o.received_by, o.created_at, o.completion_date)

	if !validateTransaction(err, tx) {
		fmt.Println("Successfully created new order")
	}
}

func (i *Order_detail) insertDetail(tx *sql.Tx) {
	statment := "INSERT INTO order_detail (order_detail_id,order_id,service_id,qty) VALUES ($1,$2,$3,$4);"

	_, err := tx.Exec(statment, i.order_detail_id, i.order_id, i.service_id, i.qty)

	if !validateTransaction(err, tx) {
		fmt.Println("Successfully created new order detail")
	}

}

func (o *Order) deleteOrder(tx *sql.Tx) {
	sqlStatment := `DELETE FROM "order"  WHERE order_id = $1`
	_, err := tx.Exec(sqlStatment, o.order_id)
	if err != nil {
		fmt.Println("Error delete order_id ", o.order_id, err)

		tx.Rollback()
	}
}

func (o *Order) deleteDetail(tx *sql.Tx) {
	sqlStatment := `DELETE FROM order_detail  WHERE order_id = $1`
	_, err := tx.Exec(sqlStatment, o.order_id)

	if err != nil {
		fmt.Println("Error delete order_detail ", err)
		tx.Rollback()
		return
	}
}
